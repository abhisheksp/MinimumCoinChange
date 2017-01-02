package minimumcoinchange

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolveForFive(t *testing.T) {
	coins := []int{1, 2, 3}
	n := 5
	minCoins := solve(n, coins)

	assert.Equal(t, 2, minCoins)
}

func TestSolveForSeven(t *testing.T) {
	coins := []int{1, 2, 3}
	n := 7
	minCoins := solve(n, coins)

	assert.Equal(t, 3, minCoins)
}

func TestSolveForEleven(t *testing.T) {
	coins := []int{9, 6, 5, 1}
	n := 11
	minCoins := solve(n, coins)

	assert.Equal(t, 2, minCoins)
}

func TestMinFirstParam(t *testing.T) {
	assert.Equal(t, 3, min(3, 4))
}

func TestMinSecondParam(t *testing.T) {
	assert.Equal(t, 3, min(4, 3))
}
