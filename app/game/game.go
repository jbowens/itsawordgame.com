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
	moment := time.Now()

	state := State{
		ID:        uuid.New(),
		CreatedAt: moment,
		UpdatedAt: moment,
		StartedAt: moment.Add(3 * time.Second),
		EndedAt:   moment.Add(3*time.Second + 2*time.Minute),
		Board:     NewBoard(defaultBoardWidth, defaultBoardHeight),
	}
	return &state
}

// AddPlayer adds a player to the game with the provided id and name.
func (s *State) AddPlayer(id string, name string) {
	s.Players[id] = Player{
		ID:   id,
		Name: name,
	}
}

// RemovePlayer removes the player with the provided id.
func (s *State) RemovePlayer(id string) {
	delete(s.Players, id)
}
