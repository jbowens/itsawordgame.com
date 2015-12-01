package game

import (
	"sort"
	"sync"

	"github.com/jbowens/dictionary"
)

var (
	dictionaryPrefixTree *dictionary.PrefixTree
)

func init() {
	d, err := dictionary.Default()
	if err != nil {
		panic(err)
	}

	d = dictionary.Filter(d, func(w string) bool { return len(w) >= 3 })
	dictionaryPrefixTree = dictionary.BuildPrefixTree(d)
}

// Solution encapsulates the solution to a board.
type Solution struct {
	root trie
}

type Answer struct {
	Path []string
	Word string
}

// CountDistinctPaths returns the number of distinct paths that yield words.
// This may be greater than the number of words if there are multiple paths
// yielding the same word.
func (s Solution) CountDistinctPaths() int {
	return s.root.Count()
}

// Words returns all the words in the board.
func (s Solution) Words() []string {
	words := s.root.Words()
	sort.Strings(words)
	return words
}

type trie struct {
	Valid bool
	Word  string
	Next  map[string]*trie
}

func (t *trie) Words() (words []string) {
	if t.Valid {
		words = append(words, t.Word)
	}
	for _, next := range t.Next {
		words = append(words, next.Words()...)
	}
	return words
}

func (t *trie) Count() (count int) {
	if t.Valid {
		count = 1
	}

	for _, t := range t.Next {
		count = count + t.Count()
	}
	return count
}

func (t *trie) Insert(word string, path []string) {
	if len(path) == 0 {
		t.Valid = true
		t.Word = word
		return
	}

	if _, ok := t.Next[path[0]]; !ok {
		t.Next[path[0]] = &trie{
			Next: map[string]*trie{},
		}
	}
	t.Next[path[0]].Insert(word, path[1:])
}

// FindSolution finds all words paths for the provided board.
func FindSolution(board Board) Solution {
	s := Solution{
		root: trie{
			Next: make(map[string]*trie),
		},
	}

	// Recursively traverse the board.
	ch := make(chan Answer)
	var wg sync.WaitGroup
	for r := 0; r < board.Height; r++ {
		for c := 0; c < board.Width; c++ {
			wg.Add(1)
			go func(r, c int) {
				generateSolution(ch, board, dictionaryPrefixTree, []string{}, make(map[string]struct{}), "", c, r)
				wg.Done()
			}(r, c)
		}
	}
	go func() {
		wg.Wait()
		close(ch)
	}()

	for answer := range ch {
		s.root.Insert(answer.Word, answer.Path)
	}
	return s
}

func generateSolution(answers chan Answer, board Board, prefixTree *dictionary.PrefixTree, cells []string, visited map[string]struct{}, word string, x int, y int) {
	if !board.WithinBounds(x, y) {
		// Out of bounds
		return
	}

	c := board.Get(x, y)
	if _, ok := visited[c.ID]; ok {
		// Already visited in this path
		return
	}

	word = word + string(c.Letter)
	prefixTree = prefixTree.Next(c.Letter)
	if prefixTree == nil {
		// There are no words with the current prefix.
		return
	}

	cells = append(cells, c.ID)
	if prefixTree.Valid {
		answers <- Answer{
			Path: cells,
			Word: word,
		}
	}

	visited[c.ID] = struct{}{}
	generateSolution(answers, board, prefixTree, cells, visited, word, x-1, y-1)
	generateSolution(answers, board, prefixTree, cells, visited, word, x-1, y)
	generateSolution(answers, board, prefixTree, cells, visited, word, x-1, y+1)
	generateSolution(answers, board, prefixTree, cells, visited, word, x, y-1)
	generateSolution(answers, board, prefixTree, cells, visited, word, x, y+1)
	generateSolution(answers, board, prefixTree, cells, visited, word, x+1, y-1)
	generateSolution(answers, board, prefixTree, cells, visited, word, x+1, y)
	generateSolution(answers, board, prefixTree, cells, visited, word, x+1, y+1)
	delete(visited, c.ID)
}
