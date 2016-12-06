package blokus

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsSolvableCorners(t *testing.T) {
	b1 := NewBoard(5)

	// should fail
	b1.Value = [][]int{
		{0, 1, 0, 0, 0},
		{1, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
		{1, 1, 0, 0, 0},
		{0, 1, 0, 0, 0},
	}

	b2 := NewBoard(5)

	// should fail
	b2.Value = [][]int{
		{0, 0, 0, 1, 0},
		{0, 0, 0, 0, 1},
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 1},
		{0, 0, 0, 1, 0},
	}
	b3 := NewBoard(5)

	// should fail
	b3.Value = [][]int{
		{0, 0, 0, 0, 0},
		{1, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
		{1, 1, 0, 0, 0},
		{0, 1, 0, 0, 0},
	}

	b4 := NewBoard(5)

	// should fail
	b4.Value = [][]int{
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 1},
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 1},
		{0, 0, 0, 1, 0},
	}

	assert.False(t, b1.IsSolvable(), "b1 Two Corners")
	assert.False(t, b2.IsSolvable(), "b2 Two Corners")
	assert.True(t, b3.IsSolvable(), "b3 False case")
	assert.True(t, b4.IsSolvable(), "b4 False case")
}

func TestIsSolvableOtherConditions(t *testing.T) {
	b1 := NewBoard(5)

	// should fail
	b1.Value = [][]int{
		{0, 1, 0, 0, 0},
		{1, 0, 0, 0, 0},
		{0, 1, 0, 0, 0},
		{1, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
	}

	b2 := NewBoard(5)

	// should fail
	b2.Value = [][]int{
		{0, 1, 0, 1, 0},
		{1, 0, 1, 0, 0},
		{0, 1, 0, 0, 0},
		{1, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
	}
	b3 := NewBoard(5)

	// should pass
	b3.Value = [][]int{
		{0, 0, 0, 0, 0},
		{1, 0, 0, 0, 1},
		{0, 1, 0, 1, 0},
		{1, 0, 0, 1, 0},
		{0, 0, 0, 0, 1},
	}

	b4 := NewBoard(5)

	// should fail
	b4.Value = [][]int{
		{0, 0, 0, 0, 0},
		{1, 0, 1, 0, 1},
		{0, 1, 0, 1, 0},
		{1, 0, 1, 1, 0},
		{0, 0, 0, 0, 1},
	}

	assert.False(t, b1.IsSolvable(), "otherConditions#b1 Mixed Edges and corners")
	assert.False(t, b2.IsSolvable(), "otherConditions#b2 Mixed Edges and corners")
	assert.True(t, b3.IsSolvable(), "otherConditions#b3 mixed Edges and middle, should be true")
	assert.False(t, b4.IsSolvable(), "otherConditions#b4 Center and edges, should be false")
}

func TestIsSolvableFailLarge(t *testing.T) {
	b1 := NewBoard(9)

	// should fail
	b1.Value = [][]int{
		{1, 0, 0, 1, 1, 1, 0, 0, 0},
		{1, 1, 1, 1, 1, 0, 0, 1, 1},
		{1, 1, 1, 0, 1, 1, 1, 1, 1},
		{1, 1, 1, 1, 0, 0, 1, 0, 0},
		{1, 1, 1, 1, 1, 0, 1, 0, 0},
		{1, 1, 1, 1, 1, 0, 1, 1, 1},
		{0, 1, 1, 1, 1, 1, 1, 0, 0},
		{1, 1, 1, 1, 0, 1, 1, 0, 0},
		{1, 0, 0, 0, 0, 0, 0, 0, 0},
	}

	assert.False(t, b1.IsSolvable(), "should fail, real case")
}
