package game

import (
	"time"

	"code.google.com/p/go-uuid/uuid"
)

const (
	defaultBoardWidth  = 5
	defaultBoardHeight = 5
)

// State encapsulates the state of a game.
type State struct {
	ID        string            `json:"id"`
	CreatedAt time.Time         `json:"created_at"`
	UpdatedAt time.Time         `json:"updated_at"`
	StartedAt time.Time         `json:"started_at"`
	EndedAt   time.Time         `json:"ended_at"`
	Board     Board             `json:"board"`
	Players   map[string]Player `json:"players"`
}

// New constructs a new game.
func New() *State {
	state := State{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Board:     NewBoard(defaultBoardWidth, defaultBoardHeight),
	}

	return &state
}
