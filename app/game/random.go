package game

import (
	"math/rand"
	"time"

	"code.google.com/p/go-uuid/uuid"
)

var (
	letterFrequencies = []struct {
		Letter             rune
		PercentThousandths int
	}{
		{'A', 8167},
		{'B', 1492},
		{'C', 2782},
		{'D', 4253},
		{'E', 12702},
		{'F', 2228},
		{'G', 2015},
		{'H', 6094},
		{'I', 6966},
		{'J', 153},
		{'K', 772},
		{'L', 4025},
		{'M', 2406},
		{'N', 6749},
		{'O', 7507},
		{'P', 1929},
		{'Q', 95},
		{'R', 5987},
		{'S', 6327},
		{'T', 9056},
		{'U', 2758},
		{'V', 978},
		{'W', 2361},
		{'X', 150},
		{'Y', 1974},
		{'Z', 74},
	}
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// randomLetter returns a random letter, according to the letter frequencies in the
// English language.
func randomLetter() rune {
	n := rand.Intn(100000)

	var sum int
	for _, lf := range letterFrequencies {
		sum = sum + lf.PercentThousandths

		if n < sum {
			return lf.Letter
		}
	}

	return letterFrequencies[len(letterFrequencies)-1].Letter
}

// randomCell returns a new random cell with a uuid ID and a random letter.
func randomCell() Cell {
	return Cell{
		ID:     uuid.New(),
		Letter: randomLetter(),
	}
}
