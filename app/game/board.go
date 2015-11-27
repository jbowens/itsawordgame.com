package game

import "bytes"

// Cell represents a single cell on the the board.
type Cell struct {
	ID     string `json:"id"`
	Letter rune   `json:"letter"`
}

// Board represents an entire itsawordgame.com game board.
type Board struct {
	Width  int                 `json:"width"`
	Height int                 `json:"height"`
	Cells  []Cell              `json:"cells"`
	idMap  map[string]Location `json:"-"`
}

// Location represents a spot on the board that a cell can occupy.
type Location struct {
	X int `json:"x"`
	Y int `json:"y"`
}

// IsAdjacent returns whether the this location and the provided location are adjacent.
func (l Location) IsAdjacent(o Location) bool {
	if l.X != o.X-1 && l.X != o.X && l.X != o.X+1 {
		return false
	}
	if l.Y != o.Y-1 && l.Y != o.Y && l.Y != o.Y+1 {
		return false
	}
	return true
}

// Get retrieves a cell by its location.
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
		idMap:  map[string]Location{},
	}

	for i := 0; i < width*height; i++ {
		c := randomCell()
		b.Cells = append(b.Cells, c)

		b.idMap[c.ID] = Location{X: i % width, Y: i / width}
	}
	return b
}
