package app

import (
	"strings"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/jbowens/itsawordgame.com/app/game"
)

type State string

const (
	WaitingForPlayers State = "waiting_for_players"
	ActiveGame        State = "active"
	GameReview        State = "review"
)

// gamekeeper manages the active game states and handles the mapping from
// websocket transport into game interactions.
type gamekeeper struct {
	// Game state management
	CurrentState State
	ActiveGame   *game.State
	ticker       *time.Ticker

	// Client requests
	incomingMessages chan incomingMessage

	// Client management
	ConnectingClients    chan *client
	DisconnectingClients chan *client
	ActiveClients        map[string]*client
}

func (gk *gamekeeper) init() {
	gk.incomingMessages = make(chan incomingMessage)
	gk.ConnectingClients = make(chan *client)
	gk.DisconnectingClients = make(chan *client)
	gk.ActiveClients = make(map[string]*client)

	// Always start off in the WaitingForPlayers state.
	gk.CurrentState = WaitingForPlayers
	gk.ticker = time.NewTicker(time.Second)

	go gk.manage()
}

func (gk *gamekeeper) manage() {
	for {
		select {
		// Handle connecting clients
		case c := <-gk.ConnectingClients:
			gk.ActiveClients[c.id] = c

			// Start the game if we were waiting for a player.
			if gk.CurrentState == WaitingForPlayers {
				gk.start()
			} else if gk.CurrentState == ActiveGame {
				gk.ActiveGame.AddPlayer(c.id, c.name)
			}
		// Handle disconnecting clients
		case c := <-gk.DisconnectingClients:
			if _, ok := gk.ActiveClients[c.id]; !ok {
				log.Errorf("Game keeper doesn't know about client %s", c.id)
				continue
			}

			// Remove the departing client and associated player.
			delete(gk.ActiveClients, c.id)
			if gk.CurrentState == ActiveGame {
				gk.ActiveGame.RemovePlayer(c.id)
			}

			// If there are no longer any active players, move back into the
			// WaitingFor
			if len(gk.ActiveClients) == 0 {
				gk.transition(WaitingForPlayers)
			}
		// Game maintenance ticker
		case t := <-gk.ticker.C:
			// TODO(jackson): Do something cleaner than a constantly ticking ticker.
			if gk.CurrentState == ActiveGame && t.After(gk.ActiveGame.EndedAt) {
				// The game has ended. We should move into the game review state.
				review := ServerMessage{
					MessageType: "game_review",
					ServerTime:  t,
					Game:        gk.ActiveGame,
				}
				for _, c := range gk.ActiveClients {
					c.output <- review
				}

				gk.transition(GameReview)
			} else if gk.CurrentState == GameReview && t.After(gk.ActiveGame.EndedAt.Add(20*time.Second)) {
				// Game review has ended
				gk.start()
			}
		// Incoming messages from clients
		case msg := <-gk.incomingMessages:
			c, ok := gk.ActiveClients[msg.client.id]
			if !ok {
				log.Warningf("Discarding a message from client `%s` because that client already disconnected.", msg.client.id)
				continue
			}

			player, ok := gk.ActiveGame.Player(msg.client.id)
			if !ok {
				log.Errorf("Player `%s` is not in game `%s`", msg.client.id, gk.ActiveGame.ID)
				continue
			}

			if gk.CurrentState != ActiveGame {
				log.Warningf("Discarding a message from client `%s` because there is no active game.", msg.client.id)
				continue
			}

			if msg.MessageType == "cell_hover" {
				scoreEvents := player.Cell(gk.ActiveGame.Solution, msg.CellID)
				if len(scoreEvents) > 0 {
					c.output <- ServerMessage{
						MessageType: "score_event",
						ServerTime:  time.Now(),
						ScoreEvents: scoreEvents,
					}
				}
			}
		}
	}
}

func (gk *gamekeeper) start() {
	gk.ActiveGame = game.New()
	for _, c := range gk.ActiveClients {
		gk.ActiveGame.AddPlayer(c.id, c.name)
	}

	words := gk.ActiveGame.Solution.Words()
	log.Infof("New game `%s` with %v words: %s", gk.ActiveGame.ID, len(words), strings.Join(words, ", "))

	// Announce the new game to all active clients. The game won't actually start until
	// a few seconds later.
	announce := ServerMessage{
		MessageType: "announce_game",
		ServerTime:  time.Now(),
		Game:        gk.ActiveGame,
	}
	for _, c := range gk.ActiveClients {
		c.output <- announce
	}

	gk.transition(ActiveGame)
}

func (gk *gamekeeper) transition(state State) {
	log.Infof("Transitioning from state `%s` to state `%s`", gk.CurrentState, state)
	gk.CurrentState = state
}
