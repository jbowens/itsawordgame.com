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
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Board     Board     `json:"board"`
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
