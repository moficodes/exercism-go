package robotname

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// Robot has a name
type Robot struct {
	name string
}

// Total number of combinations for letter is 26^2 = 676
// Total number of combinations for digits is 10^3 = 1000
// combinations formula is n^m
const maxCombinations = 26 * 26 * 10 * 10 * 10

// Seen - global var to remember used names
var Seen = map[string]struct{}{"": struct{}{}}

// Name generates unique robot's name
func (r *Robot) Name() (string, error) {
	if r.name != "" {
		return r.name, nil
	}
	if len(Seen) >= maxCombinations {
		return "", errors.New("names limit reached")
	}
	for {
		if _, ok := Seen[r.name]; ok {
			r.name = generateName()
		} else {
			Seen[r.name] = struct{}{}
			break
		}
	}

	return r.name, nil
}

func generateName() string {
	rand.Seed(time.Now().UTC().UnixNano())
	return fmt.Sprintf("%s%s%d%d%d", randomLetter(), randomLetter(), randomDigit(), randomDigit(), randomDigit())
}

// Reset - resets robot's name
func (r *Robot) Reset() {
	r.name = ""
}

func randomLetter() string {
	return string('A' + rand.Intn(26))
}

func randomDigit() int {
	return rand.Intn(10)
}
