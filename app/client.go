package app

import (
	"time"

	"code.google.com/p/go-uuid/uuid"

	log "github.com/Sirupsen/logrus"
	"github.com/gorilla/websocket"
)

const (
	outputChannelBuffer = 5
	writeWait           = 5 * time.Second
	pingFrequency       = 10 * time.Second
	maxMessageSize      = 1024 * 1024
)

type incomingMessage struct {
	ClientMessage
	client *client
}

func newClient(remoteAddr string, conn *websocket.Conn, incoming chan incomingMessage) *client {
	c := &client{
		id:         uuid.New(),
		remoteAddr: remoteAddr,
		pingTicker: time.NewTicker(pingFrequency),
		output:     make(chan ServerMessage, outputChannelBuffer),
		input:      incoming,
		conn:       conn,
	}

	log.Infof("Creating new client `%s` from %s", c.id, c.remoteAddr)

	go c.writeLoop()
	go c.readLoop()

	return c
}

type client struct {
	id         string
	remoteAddr string
	pingTicker *time.Ticker
	output     chan ServerMessage
	input      chan incomingMessage
	conn       *websocket.Conn
}

func (c *client) ping() error {
	log.Infof("Pinging client %s", c.id)
	c.conn.SetWriteDeadline(time.Now().Add(writeWait))
	return c.conn.WriteMessage(websocket.PingMessage, []byte{})
}

func (c *client) closeConnection() error {
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
		c.pingTicker.Stop()
		c.conn.Close()
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
				c.closeConnection()
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
		var clientMessage ClientMessage
		if err := c.conn.ReadJSON(&clientMessage); err != nil {
			c.conn.Close()
			break
		}

		c.input <- incomingMessage{
			ClientMessage: clientMessage,
			client:        c,
		}
	}
}
