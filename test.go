package blokus

import (
	"testing"

	"fmt"

	"github.com/stretchr/testify/assert"
)

func TestIsSolvableCorners(t *testing.T) {
	fmt.Println("In the test")
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

	assert.False(t, b1.IsSolvable(), "b1 Two Corners")
	assert.False(t, b2.IsSolvable(), "b2 Two Corners")
}
