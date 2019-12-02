package scorekeeper_test

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/scorekeeper"
)

// get the score of a player at start returns 0
func TestGetInitialScore(t *testing.T) {
	scorer := scorekeeper.NewScorer()
	assert.Equal(t, 0, scorer.Get("Sandy"))
}

// add 10 to the score of a player and then get returns 10
func TestAddPoints(t *testing.T) {
	scorer := scorekeeper.NewScorer()
	scorer.Add("Sandy", 10)
	assert.Equal(t, 10, scorer.Get("Sandy"))
}

// add 10 to the score of a player and then another 10 and get returns 20
func TestAddPointsAgain(t *testing.T) {
	scorer := scorekeeper.NewScorer()
	scorer.Add("Sandy", 10)
	scorer.Add("Sandy", 10)
	assert.Equal(t, 20, scorer.Get("Sandy"))
}

// subtract 5 from a score of 10 to get 5
func TestSubstractPoints(t *testing.T) {
	scorer := scorekeeper.NewScorer()
	scorer.Add("Sandy", 10)
	scorer.Substract("Sandy", 5)
	assert.Equal(t, 5, scorer.Get("Sandy"))
}

// subtract 5 from a score of 10 to get 5 and then subtract 5 again to get to 0
func TestSubstractPointsAgain(t *testing.T) {
	scorer := scorekeeper.NewScorer()
	scorer.Add("Sandy", 10)
	scorer.Substract("Sandy", 5)
	scorer.Substract("Sandy", 5)
	assert.Equal(t, 0, scorer.Get("Sandy"))
}

// subtract 5 from a score of 0 to get 0
func TestSubstractPointsFromZeroReturnsZero(t *testing.T) {
	scorer := scorekeeper.NewScorer()
	scorer.Substract("Sandy", 5)
	assert.Equal(t, 0, scorer.Get("Sandy"))
}

// subtract 5 from a score of 2 to get 0
func TestSubstractPointsGreaterThanCurrentScoreReturnsZero(t *testing.T) {
	scorer := scorekeeper.NewScorer()
	scorer.Add("Sandy", 2)
	scorer.Substract("Sandy", 5)
	assert.Equal(t, 0, scorer.Get("Sandy"))
}

// add 5 from a scorer A and add 10 for scorer B and then get A's and B's scores
func TestAddPointsMultipleScorers(t *testing.T) {
	scorerA := scorekeeper.NewScorer()
	scorerB := scorekeeper.NewScorer()
	scorerA.Add("Sandy", 5)
	scorerB.Add("George", 2)
	assert.Equal(t, 5, scorerA.Get("Sandy"))
	assert.Equal(t, 2, scorerB.Get("George"))
}

// add 5 from a scorer A and add 10 for scorer B and then get A's and B's scores
func TestSubstractPointsMultipleScorers(t *testing.T) {
	scorerA := scorekeeper.NewScorer()
	scorerB := scorekeeper.NewScorer()
	scorerA.Add("Sandy", 5)
	scorerB.Add("George", 2)
	scorerA.Substract("Sandy", 2)
	scorerB.Substract("George", 2)
	assert.Equal(t, 3, scorerA.Get("Sandy"))
	assert.Equal(t, 0, scorerB.Get("George"))
}