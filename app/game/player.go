package game

// Player represents an individual player in a particular game.
type Player struct {
	ID      string      `json:"id"`
	Name    string      `json:"name"`
	Scores  []Score     `json:"scores"`
	Current pathBuilder `json:"-"`
}

// Points totals all this player's points.
func (p *Player) Points() (total int) {
	for _, s := range p.Scores {
		total = total + s.Points
	}
	return total
}

// Cell adds the provided cell to the player's path and processes any events
// stemming from it.
func (p *Player) Cell(id string) error {
	return p.Current.Append(id)
}
