package game

import "bytes"

// Cell represents a single cell on the the board.
type Cell struct {
	ID     string
	Letter rune
}

// Board represents an entire itsawordgame.com game board.
type Board struct {
	Width, Height int
	Cells         []Cell
	idMap         map[string]coord
}

type coord struct {
	X, Y int
}

// Get retrieves a cell by its coordinates.
func (b Board) Get(x, y int) Cell {
	return b.Cells[y*b.Height+x]
}

// String returns a string representation of the board.
func (b Board) String() string {
	var buf bytes.Buffer

	for r := 0; r < b.Height; r++ {
		rowCells := b.Cells[r*b.Width : (r+1)*b.Width]
		for _, c := range rowCells {
			buf.WriteRune(c.Letter)
			buf.WriteRune('\t')
		}
		buf.WriteRune('\n')
		buf.WriteRune('\n')
	}

	return buf.String()
}

// NewBoard generates a new random board.
func NewBoard(width, height int) Board {
	b := Board{
		Width:  width,
		Height: height,
		Cells:  make([]Cell, 0, width*height),
		idMap:  map[string]coord{},
	}

	for i := 0; i < width*height; i++ {
		c := randomCell()
		b.Cells = append(b.Cells, c)

		b.idMap[c.ID] = coord{X: i % width, Y: i / width}
	}
	return b
}
