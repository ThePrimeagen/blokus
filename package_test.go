package blokus

import (
	"testing"

	"fmt"

	"github.com/stretchr/testify/assert"
)

func TestIsSolvableCorners(t *testing.T) {
	b := NewBoard(5)
	p1 := &Piece{2, 2,
		false,
		false,
		[][]int{
			{1, 1},
			{0, 1},
		},
	}
	res1 := b.Add(p1, 3, 0)

	assert.True(t, res1, "Expected the add to be successful")
	assert.True(t, b.IsSolvable(), "Graph is still solvable")
	assert.Equal(t, 1, b.ones, "ones should be equal 1")

	p2 := &Piece{1, 1,
		true,
		false,
		[][]int{
			{1},
		},
	}

	res2 := b.Add(p2, 0, 0)
	assert.True(t, res2, "Expected the add to be successful")
	assert.False(t, b.IsSolvable(), "Graph is not solvable")
	assert.Equal(t, 1, b.ones, "ones should be equal 1")
	assert.True(t, b.hasOne, "Should have hasOne to be true")

	after2 := [][]int{
		{12, 2, 1, 1, 2},
		{2, 0, 0, 0, 1},
		{2, 1, 0, 0, 1},
		{12, 12, 1, 0, 1},
		{4, 12, 2, 1, 2},
	}

	assert.Equal(t, after2, b.Value, fmt.Sprintf("Expected the Board to be equal %v", b))

	b.Remove(p2, 0, 0)
	after3 := [][]int{
		{2, 1, 1, 1, 2},
		{1, 0, 0, 0, 1},
		{2, 1, 0, 0, 1},
		{12, 12, 1, 0, 1},
		{4, 12, 2, 1, 2},
	}
	assert.Equal(t, after3, b.Value, fmt.Sprintf("Expected the Board to be equal %v", b))
	assert.True(t, b.IsSolvable(), "Graph is solvable again")
	assert.Equal(t, 1, b.ones, "ones should be equal 1")
	assert.False(t, b.hasOne, "We do not have a one anymore")

	b.Remove(p1, 3, 0)
	after4 := [][]int{
		{2, 1, 1, 1, 2},
		{1, 0, 0, 0, 1},
		{1, 0, 0, 0, 1},
		{1, 0, 0, 0, 1},
		{2, 1, 1, 1, 2},
	}
	assert.Equal(t, after4, b.Value, fmt.Sprintf("Expected the Board to be equal %v", b))
	assert.True(t, b.IsSolvable(), "Graph is blank, so should be solvable.")
	assert.Equal(t, 0, b.ones, "ones should be equal 0")
	assert.False(t, b.hasOne, "We do not have a one anymore")

	p4 := &Piece{4, 2,
		false,
		false,
		[][]int{
			{1, 1},
			{1, 0},
			{1, 0},
			{1, 1},
		},
	}

	b.Add(p4, 0, 3)
	after5 := [][]int{
		{2, 1, 2, 13, 13},
		{1, 0, 1, 12, 3},
		{1, 0, 1, 12, 3},
		{1, 0, 1, 12, 12},
		{2, 1, 1, 2, 3},
	}
	assert.Equal(t, after5, b.Value, fmt.Sprintf("Expected the Board to be equal %v", b))
	assert.True(t, b.IsSolvable(), "Graph has p4 only")
	assert.Equal(t, 0, b.ones, "ones should be equal 0")
	assert.Equal(t, 1, b.twos, "twos should be equal 1")
	assert.False(t, b.hasOne, "hasOne == false")
	assert.False(t, b.hasTwo, "hasTwo == false")

	p3 := &Piece{2, 1,
		false,
		true,
		[][]int{
			{1},
			{1},
		},
	}
	b.Add(p3, 0, 0)
	after6 := [][]int{
		{13, 2, 2, 13, 13},
		{12, 1, 1, 12, 3},
		{2, 0, 1, 12, 3},
		{1, 0, 1, 12, 12},
		{2, 1, 1, 2, 3},
	}
	assert.Equal(t, after6, b.Value, fmt.Sprintf("Expected the Board to be equal %v", b))
	assert.False(t, b.IsSolvable(), "Graph has p4 and p3")
	assert.Equal(t, 0, b.ones, "ones should be equal 0")
	assert.Equal(t, 1, b.twos, "twos should be equal 1")
	assert.False(t, b.hasOne, "hasOne == false")
	assert.True(t, b.hasTwo, "hasTwo == true")

	b.Remove(p3, 0, 0)
	b.Add(p3, 1, 4)
	after8 := [][]int{
		{2, 1, 2, 13, 14},
		{1, 0, 1, 13, 14},
		{1, 0, 1, 13, 14},
		{1, 0, 1, 12, 13},
		{2, 1, 1, 2, 3},
	}
	assert.Equal(t, after8, b.Value, fmt.Sprintf("Expected the Board to be equal %v", b))
	assert.True(t, b.IsSolvable(), "Graph has p4 and p3")
	assert.Equal(t, 0, b.ones, "ones should be equal 0")
	assert.Equal(t, 0, b.twos, "twos should be equal 1")
	assert.False(t, b.hasOne, "hasOne == false")
	assert.True(t, b.hasTwo, "hasTwo == true")
}
