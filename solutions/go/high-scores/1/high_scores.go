package highscores

import (
	"cmp"
	"math"
	"slices"
)

type HighScores struct {
	topThree     []int
	personalBest int
	latest       int
	scores       []int
}

// NewHighScores returns a new HighScores object.
func NewHighScores(scores []int) *HighScores {
	sortedScores := slices.Clone(scores)
	slices.SortFunc(sortedScores, func(a, b int) int {
		return cmp.Compare(b, a)
	})

	return &HighScores{
		topThree:     sortedScores[:int(math.Min(float64(len(scores)), 3))],
		personalBest: sortedScores[0],
		latest:       scores[len(scores)-1],
		scores:       scores,
	}
}

// Scores returns all the scores.
func (s *HighScores) Scores() []int {
	return s.scores
}

// Latest returns the latest (last) score.
func (s *HighScores) Latest() int {
	return s.latest
}

// PersonalBest returns the best (highest) score.
func (s *HighScores) PersonalBest() int {
	return s.personalBest
}

// TopThree returns the top three scores.
func (s *HighScores) TopThree() []int {
	return s.topThree
}
