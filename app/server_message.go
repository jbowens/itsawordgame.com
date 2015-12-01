package app

import (
	"time"

	"github.com/jbowens/itsawordgame.com/app/game"
)

// ServerMessage defines the schema for json messages sent from the server
// to the client.
type ServerMessage struct {
	MessageType string       `json:"message_type"`
	ServerTime  time.Time    `json:"server_time"`
	Game        *game.State  `json:"game,omitempty"`
	ScoreEvents []game.Score `json:"score_events,omitempty"`
}
