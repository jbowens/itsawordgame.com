package game

import "time"

// Score represents an instance when a player earned points.
type Score struct {
	ScoredAt time.Time `json:"scored_at"`
	Word     string    `json:"word"`
	Points   int       `json:"points"`
}
