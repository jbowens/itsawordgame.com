package game

import "time"

// Player represents an individual player in a particular game.
type Player struct {
	ID       string              `json:"id"`
	Name     string              `json:"name"`
	Scores   []Score             `json:"scores"`
	Words    map[string]struct{} `json:"-"`
	Current  []*trie             `json:"-"`
	LastCell string              `json:"-"`
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
func (p *Player) Cell(solution Solution, id string) []Score {
	if p.LastCell == id {
		// Ignore, don't invalidate the path they've built up. It's really common to
		// hover over the same cell twice in a row accidentally.
		return nil
	}
	p.LastCell = id

	var scores []Score
	newNodes := make([]*trie, 0, len(p.Current))
	if p.Words == nil {
		p.Words = map[string]struct{}{}
	}

	current := append(p.Current, &solution.root)
	for _, t := range current {
		next, ok := t.Next[id]
		if !ok {
			continue
		}

		newNodes = append(newNodes, next)
		if next.Valid {
			if _, ok := p.Words[next.Word]; !ok {
				scores = append(scores, Score{
					ScoredAt: time.Now(),
					Word:     next.Word,
					Points:   len(next.Word),
				})
				p.Words[next.Word] = struct{}{}
			}
		}
	}
	p.Current = newNodes
	return scores
}
