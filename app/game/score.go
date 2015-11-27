package game

import "time"

// Score represents an instance when a player earned points.
type Score struct {
	ScoredAt time.Time `json:"scored_at"`
	Path     Path      `json:"path"`
	Points   int       `json:"points"`
}
