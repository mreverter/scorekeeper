package scorekeeper_test

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/scorekeeper"
)

func TestGetInitialScore(t *testing.T) {
	scorer := scorekeeper.NewScorer()
	assert.Equal(t, 0, scorer.Get("Home"))
}

func TestAddPointsToScore(t *testing.T) {
	scorer := scorekeeper.NewScorer()
	scorer.Add(5, "Home")
	assert.Equal(t, 5, scorer.Get("Home"))
}

func TestAddPointsMultipleTimesToScore(t *testing.T) {
	scorer := scorekeeper.NewScorer()
	scorer.Add(5, "Home")
	scorer.Add(5, "Home")
	assert.Equal(t, 10, scorer.Get("Home"))
}

func TestGetInitialScoresForMultiplePlayers(t *testing.T) {
	scorer := scorekeeper.NewScorer()
	assert.Equal(t, 0, scorer.Get("Home"))
	assert.Equal(t, 0, scorer.Get("Away"))
}

func TestAddPointsToOnePlayerAndNotOthers(t *testing.T) {
	scorer := scorekeeper.NewScorer()
	scorer.Add(10, "Home")
	assert.Equal(t, 10, scorer.Get("Home"))
	assert.Equal(t, 0, scorer.Get("Away"))
}
