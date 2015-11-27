package game

import (
	"crypto/sha1"
	"fmt"
	"io"
)

// Path refers to a path across the board. It's stored as a slice of the IDs of the
// cells in the order that they're selected.
type Path []string

// Hash returns a hash of the entire path, serving as a unique identifier for this
// path.
func (p Path) Hash() string {
	h := sha1.New()
	for _, id := range p {
		io.WriteString(h, id)
	}
	return fmt.Sprintf("%x", h.Sum(nil))
}

// Valid returns true if the path is a valid path on the provided board.
func (p Path) Valid(board Board) bool {
	ids := map[string]struct{}{}

	var prev *Location

	for _, id := range p {
		// Check that the cell isn't duplicated.
		if _, ok := ids[id]; ok {
			return false
		}
		ids[id] = struct{}{}

		// Check that the cell actually exists on this board.
		l, ok := board.idMap[id]
		if !ok {
			return false
		}

		// Check that the cell is adjacent to the previous cell.
		if prev != nil && !prev.IsAdjacent(l) {
			return false
		}
		prev = &l
	}

	return true
}
