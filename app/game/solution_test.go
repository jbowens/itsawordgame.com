package game

import "testing"

func TestFindSolution(t *testing.T) {
	b := Board{
		Width:  5,
		Height: 4,
		Cells: []Cell{
			{"0", 'E'},
			{"1", 'A'},
			{"2", 'T'},
			{"3", 'C'},
			{"4", 'C'},
			{"5", 'O'},
			{"6", 'S'},
			{"7", 'N'},
			{"8", 'I'},
			{"9", 'H'},
			{"10", 'I'},
			{"11", 'L'},
			{"12", 'E'},
			{"13", 'M'},
			{"14", 'S'},
			{"15", 'A'},
			{"16", 'W'},
			{"17", 'O'},
			{"18", 'D'},
			{"19", 'E'},
		},
		idMap: map[string]Location{
			"0":  {0, 0},
			"1":  {1, 0},
			"2":  {2, 0},
			"3":  {3, 0},
			"4":  {4, 0},
			"5":  {0, 1},
			"6":  {1, 1},
			"7":  {2, 1},
			"8":  {3, 1},
			"9":  {4, 1},
			"10": {0, 2},
			"11": {1, 2},
			"12": {2, 2},
			"13": {3, 2},
			"14": {4, 2},
			"15": {0, 3},
			"16": {1, 3},
			"17": {2, 3},
			"18": {3, 3},
			"19": {4, 3},
		},
	}

	solution := FindSolution(b)
	words := solution.Words()
	set := map[string]struct{}{}
	for _, w := range words {
		set[w] = struct{}{}
	}

	testWords := []string{"CHIME", "EAST", "MINES", "MINT", "SANE", "SHINE", "TIME"}
	for _, w := range testWords {
		if _, ok := set[w]; !ok {
			t.Errorf("Expected to find `%s` in solution but couldn't find it for board:\n\n%s", w, b)
		}
		if !dictionaryPrefixTree.Contains(w) {
			t.Errorf("Expected the dictionaryPrefixTree to contain `%s` but couldn't find it.", w)
		}
	}

	m, ok := solution.root.Next["13"] // M
	if !ok {
		t.Fatalf("root -> 13: does not exist: %+v", solution.root)
	}
	i, ok := m.Next["8"] // I
	if !ok {
		t.Fatalf("root -> 13 -> 8 does not exist: %+v", m)
	}
	n, ok := m.Next["7"] // N
	if !ok {
		t.Fatalf("root -> 13 -> 8 -> 7 does not exist: %+v", i)
	}
	mint, ok := m.Next["2"] // T
	if !ok {
		t.Fatalf("root -> 13 -> 8 -> 7 -> 2 does not exist: %+v", n)
	}
	if !mint.Valid {
		t.Fatalf("root -> 13 -> 8 -> 7 -> 2 should be marked as a valid word: %+v", mint)
	}
	if mint.Word != "MINT" {
		t.Fatalf("root -> 13 -> 8 -> 7 -> 2 should be the word 'MINT': %+v", mint)
	}
}
