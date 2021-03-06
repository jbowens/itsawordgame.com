package game

import (
	"time"

	"github.com/pborman/uuid"
)

const (
	defaultBoardWidth  = 5
	defaultBoardHeight = 4

	// boardMinimumWordCount defines the minimum number of words that should be in
	// a game board. If we randomly generate a board with fewer words than this
	// constant, we generate a new board.
	boardMinimumWordCount = 150
)

// State encapsulates the state of a game.
type State struct {
	ID        string             `json:"id"`
	CreatedAt time.Time          `json:"created_at"`
	UpdatedAt time.Time          `json:"updated_at"`
	StartedAt time.Time          `json:"started_at"`
	EndedAt   time.Time          `json:"ended_at"`
	Board     Board              `json:"board"`
	Players   map[string]*Player `json:"players"`
	Solution  Solution           `json:"-"`
}

// New constructs a new game.
func New() *State {
	moment := time.Now()

	board := NewBoard(defaultBoardWidth, defaultBoardHeight)
	solution := FindSolution(board)
	for len(solution.Words()) < boardMinimumWordCount {
		board = NewBoard(defaultBoardWidth, defaultBoardHeight)
		solution = FindSolution(board)
	}

	state := State{
		ID:        uuid.New(),
		CreatedAt: moment,
		UpdatedAt: moment,
		StartedAt: moment.Add(3 * time.Second),
		EndedAt:   moment.Add(3*time.Second + 2*time.Minute),
		Board:     board,
		Players:   make(map[string]*Player),
		Solution:  solution,
	}
	return &state
}

// Player retrieves the player with the provided id.
func (s *State) Player(id string) (*Player, bool) {
	p, ok := s.Players[id]
	return p, ok
}

// AddPlayer adds a player to the game with the provided id and name.
func (s *State) AddPlayer(id string, name string) {
	s.Players[id] = &Player{
		ID:   id,
		Name: name,
	}
}

// RemovePlayer removes the player with the provided id.
func (s *State) RemovePlayer(id string) {
	delete(s.Players, id)
}
