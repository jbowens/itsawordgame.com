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

type pathBuilder struct {
	prev  Location
	board Board
	path  Path
}

func (pb *pathBuilder) Append(newID string) error {
	l, ok := pb.board.LookupCellByID(newID)
	if !ok {
		return fmt.Errorf("Cell `%s` not found", newID)
	}

	// Check if this cell is already in the path. If it is, we need to trim the start
	// of the path to trim out the previous use of this same cell.
	for i, existingID := range pb.path {
		if existingID == newID {
			// This cell is already in our path. We need to reset the path to start just
			// after the old use of this cell. Note that we make a new copy of the array to
			// to prevent the array backing the slice from monotonically increasing forever.
			newPath := make([]string, 0, 2*len(pb.path[i+1:]))
			newPath = append(newPath, pb.path[i+1:]...)
			pb.path = newPath
		}
	}

	// If the path is empty, then we don't need to check that pb.prev is adjacent.
	if len(pb.path) == 0 {
		pb.path = []string{newID}
		pb.prev = l
		return nil
	}

	if !pb.prev.IsAdjacent(l) {
		// If the locations are not adjacent on the board, we need to forget the entire
		// existing path.
		pb.path = nil
	}

	pb.path = append(pb.path, newID)
	pb.prev = l
	return nil
}
