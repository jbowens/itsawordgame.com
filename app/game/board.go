package game

// Cell represents a single cell on the the board.
type Cell struct {
	ID     string
	Letter rune
}

// Board represents an entire itsawordgame.com game board.
type Board struct {
	Cells [][]Cell
}
