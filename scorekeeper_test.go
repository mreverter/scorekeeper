package scorekeeper_test

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/scorekeeper"
)

func TestGetInitialScore(t *testing.T) {
	scorer := scorekeeper.NewScorer()
	assert.Equal(t, 0, scorer.Get())
}

func TestAddPointsToScore(t *testing.T) {
	scorer := scorekeeper.NewScorer()
	scorer.Add(5)
	assert.Equal(t, 5, scorer.Get())
}

func TestAddPointsMultipleTimesToScore(t *testing.T) {
	scorer := scorekeeper.NewScorer()
	scorer.Add(5)
	scorer.Add(5)
	assert.Equal(t, 10, scorer.Get())
}