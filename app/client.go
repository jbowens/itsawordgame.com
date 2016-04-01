package app

import (
	"fmt"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/gorilla/websocket"
	"github.com/pborman/uuid"
)

const (
	outputChannelBuffer = 10
	writeWait           = 5 * time.Second
	pingFrequency       = 10 * time.Second
	maxMessageSize      = 1024 * 1024
)

type incomingMessage struct {
	ClientMessage
	client *client
}

func newClient(remoteAddr string, conn *websocket.Conn, incoming chan incomingMessage, closing chan *client) *client {
	c := &client{
		id:         uuid.New(),
		remoteAddr: remoteAddr,
		pingTicker: time.NewTicker(pingFrequency),
		output:     make(chan ServerMessage, outputChannelBuffer),
		input:      incoming,
		closing:    closing,
		closed:     make(chan struct{}),
		conn:       conn,
	}

	// TODO(jackson): Support joining with a custom name
	c.name = fmt.Sprintf("Guest %s", c.id[:6])

	log.Infof("Creating new client `%s` from %s", c.id, c.remoteAddr)

	go c.writeLoop()
	go c.readLoop()

	return c
}

type client struct {
	id         string
	name       string
	remoteAddr string
	pingTicker *time.Ticker
	output     chan ServerMessage
	input      chan incomingMessage
	closing    chan *client
	closed     chan struct{}
	conn       *websocket.Conn
}

func (c *client) exit(graceful bool) {
	// Non-blocking check if we're already closed.
	select {
	case <-c.closed:
		return
	default:
		log.Infof("Client `%s` shutdown started", c.id)
	}
	close(c.closed)

	if graceful {
		if err := c.sendClose(); err != nil {
			log.Errorf("error closing connection: %s", err)
		}
	}
	c.conn.Close()
	c.pingTicker.Stop()
	c.closing <- c
}

func (c *client) ping() error {
	log.Infof("Pinging client %s", c.id)
	c.conn.SetWriteDeadline(time.Now().Add(writeWait))
	return c.conn.WriteMessage(websocket.PingMessage, []byte{})
}

func (c *client) sendClose() error {
	c.conn.SetWriteDeadline(time.Now().Add(writeWait))
	return c.conn.WriteMessage(websocket.CloseMessage, []byte{})
}

func (c *client) sendJSON(obj interface{}) error {
	c.conn.SetWriteDeadline(time.Now().Add(writeWait))
	return c.conn.WriteJSON(obj)
}

func (c *client) writeLoop() {
	defer func() {
		log.Infof("Client %s exiting", c.id)
		c.exit(false)
	}()

	for {
		select {
		case <-c.pingTicker.C:
			if err := c.ping(); err != nil {
				log.Errorf("Error pinging client %s: %v", c.id, err)
				return
			}
		case msg, ok := <-c.output:
			if !ok {
				c.exit(true)
				return
			}

			if err := c.sendJSON(msg); err != nil {
				log.Errorf("Error sending message to client %s: %v", c.id, err)
				return
			}
		}
	}
}

func (c *client) readLoop() {
	defer log.Infof("Exiting read routine for client %s", c.id)

	c.conn.SetReadLimit(maxMessageSize)
	c.conn.SetReadDeadline(time.Time{})
	c.conn.SetPongHandler(func(msg string) error {
		log.Debugf("Pong from client %s", c.id)
		return nil
	})
	for {
		msg := incomingMessage{client: c}
		if err := c.conn.ReadJSON(&msg.ClientMessage); err != nil {
			c.exit(false)
			return
		}

		c.input <- msg
	}
}
