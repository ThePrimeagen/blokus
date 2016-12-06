package blokus

import "fmt"

/*
PieceGroup is great
*/
type PieceGroup struct {
	Pieces []*Piece
}

func (pG *PieceGroup) String() string {
	return fmt.Sprintf("PieceGroup has %v pieces", len(pG.Pieces))
}

/*
Piece is great
*/
type Piece struct {
	Row   int
	Col   int
	Value [][]int
}

func (p *Piece) String() string {
	outStr := ""

	for col := 0; col < p.Col; col++ {
		for row := 0; row < p.Row; row++ {
			if p.Value[row][col] == 0 {
				outStr += "   "
			} else {
				outStr += "[ ]"
			}
		}

		outStr += "\n"
	}

	return outStr
}

/*
Board is great
*/
type Board struct {
	Row   int
	Col   int
	Value [][]int
}

/*
Solved is great
*/
func (b *Board) Solved() bool {
	solved := true
	for col := 0; col < b.Col && solved; col++ {
		for row := 0; row < b.Row && solved; row++ {
			solved = b.Value[row][col] == 1
		}
	}

	return solved
}

func (b *Board) String() string {
	outStr := ""

	for col := 0; col < b.Col; col++ {
		for row := 0; row < b.Row; row++ {
			if b.Value[row][col] == 0 {
				outStr += "[ ]"
			} else {
				outStr += "[x]"
			}
		}

		outStr += "\n"
	}

	return outStr
}

/*
NewBoard is great
*/
func NewBoard(size int) *Board {
	board := Board{size, size, [][]int{}}
	rows := make([][]int, size)

	boardValue := rows[:]

	for i := range rows {
		row := make([]int, size)
		rows[i] = row[0:]
	}

	board.Value = boardValue
	return &board
}

/*
Add adds the piece to the board.  If it cannot, then the function will return
false.
*/
func (b *Board) Add(p *Piece, row, col int) bool {
	if !isValid(b, p, row, col) {
		return false
	}

	for r := 0; r < p.Row; r++ {
		for c := 0; c < p.Col; c++ {
			if p.Value[r][c] == 0 {
				continue
			}

			b.Value[row+r][col+c] = 1
		}
	}

	return true
}

/*
Remove same piece from same position
*/
func (b *Board) Remove(p *Piece, row, col int) bool {

	for r := 0; r < p.Row; r++ {
		for c := 0; c < p.Col; c++ {
			if p.Value[r][c] == 0 {
				continue
			}

			b.Value[row+r][col+c] = 0
		}
	}

	return true
}

/*
IsSolvable searches for 1s and 2s spots
*/
func (b *Board) IsSolvable() bool {
	hasOne := 0
	farR, farC := b.Row-1, b.Col-1
	farRM1, farCM1 := farR-1, farC-1

	// check for the edge conditions so that this logic does
	if b.Value[0][0] == 0 &&
		b.Value[1][0] == 1 &&
		b.Value[0][1] == 1 {
		hasOne++
	}

	if b.Value[farR][farC] == 0 &&
		b.Value[farRM1][farC] == 1 &&
		b.Value[farR][farCM1] == 1 {
		hasOne++
	}

	if b.Value[farR][0] == 0 &&
		b.Value[farRM1][0] == 1 &&
		b.Value[farR][1] == 1 {
		hasOne++
	}

	if b.Value[0][farC] == 0 &&
		b.Value[0][farCM1] == 1 &&
		b.Value[1][farC] == 1 {
		hasOne++
	}

	// top row
	for c := 1; c < farC && hasOne < 2; c++ {
		if b.Value[0][c] == 0 &&
			b.Value[0][c+1] == 1 &&
			b.Value[0][c-1] == 1 &&
			b.Value[1][c] == 1 {
			hasOne++
		}
	}

	// bottom row
	for c := 1; c < farC && hasOne < 2; c++ {
		if b.Value[farR][c] == 0 &&
			b.Value[farR][c+1] == 1 &&
			b.Value[farR][c-1] == 1 &&
			b.Value[farRM1][c] == 1 {
			hasOne++
		}
	}

	// left most col
	for r := 1; r < farR && hasOne < 2; r++ {
		if b.Value[r][0] == 0 &&
			b.Value[r-1][0] == 1 &&
			b.Value[r+1][0] == 1 &&
			b.Value[r][1] == 1 {
			hasOne++
		}
	}

	// right most col
	for r := 1; r < farR && hasOne < 2; r++ {
		if b.Value[r][farC] == 0 &&
			b.Value[r-1][farC] == 1 &&
			b.Value[r+1][farC] == 1 &&
			b.Value[r][farCM1] == 1 {
			hasOne++
		}
	}

	// The middle thing
	for r := 1; r < farR && hasOne < 2; r++ {
		for c := 1; c < farC && hasOne < 2; c++ {

			// corner cases
			if b.Value[r][c] == 0 &&
				b.Value[r+1][c] == 1 &&
				b.Value[r-1][c] == 1 &&
				b.Value[r][c+1] == 1 &&
				b.Value[r][c-1] == 1 {
				hasOne++
			}
		}
	}

	return hasOne < 2
}

/*
GetPieces is great
*/
func GetPieces() []*PieceGroup {
	return []*PieceGroup{
		oneByOne,
		oneByTwo,
		oneByThree,
		oneByFour,
		oneByFive,
		lBow,
		tumorBlock,
		leggy,
		z,
		bowBow,
		bigL,
		plus,
		l,
		bow,
		block,
		pacMan,
		t,
		oddy,
		shooty,
		longL,
		tetris,
	}
}

func isValid(b *Board, p *Piece, row, col int) bool {

	collision := false
	for r := 0; r < p.Row && !collision; r++ {
		for c := 0; c < p.Col && !collision; c++ {
			collision = b.Value[row+r][col+c] == 1
		}
	}

	return !collision

}

/*
OneByOne is great
*/
var oneByOne = &PieceGroup{
	[]*Piece{
		&Piece{
			1,
			1,
			[][]int{{1}},
		},
	},
}

/*
OneByTwo is great
*/
var oneByTwo = &PieceGroup{
	[]*Piece{
		&Piece{
			2,
			1,
			[][]int{
				{1},
				{1},
			},
		},
		&Piece{
			1,
			2,
			[][]int{{1, 1}},
		},
	},
}

/*
OneByThree is great
*/
var oneByThree = &PieceGroup{
	[]*Piece{
		&Piece{
			3,
			1,
			[][]int{
				{1},
				{1},
				{1},
			},
		},
		&Piece{
			1,
			3,
			[][]int{{1, 1, 1}},
		},
	},
}

/*
OneByFour is great
*/
var oneByFour = &PieceGroup{
	[]*Piece{
		&Piece{
			4,
			1,
			[][]int{
				{1},
				{1},
				{1},
				{1},
			},
		},
		&Piece{
			1,
			4,
			[][]int{{1, 1, 1, 1}},
		},
	},
}

/*
OneByFive is great
*/
var oneByFive = &PieceGroup{
	[]*Piece{
		&Piece{
			5,
			1,
			[][]int{
				{1},
				{1},
				{1},
				{1},
				{1},
			},
		},
		&Piece{
			1,
			5,
			[][]int{{1, 1, 1, 1, 1}},
		},
	},
}

/*
LBow is great
*/
var lBow = &PieceGroup{
	[]*Piece{
		&Piece{2, 2,
			[][]int{
				{1, 0},
				{1, 1},
			},
		},
		&Piece{2, 2,
			[][]int{
				{1, 1},
				{0, 1},
			},
		},
		&Piece{2, 2,
			[][]int{
				{0, 1},
				{1, 1},
			},
		},
		&Piece{2, 2,
			[][]int{
				{1, 1},
				{1, 0},
			},
		},
	},
}

/*
TumorBlock is great
*/
var tumorBlock = &PieceGroup{
	[]*Piece{
		&Piece{2, 3,
			[][]int{
				{1, 1, 0},
				{1, 1, 1},
			},
		},
		&Piece{2, 3,
			[][]int{
				{1, 1, 1},
				{1, 1, 0},
			},
		},
		&Piece{2, 3,
			[][]int{
				{1, 1, 1},
				{0, 1, 1},
			},
		},
		&Piece{2, 3,
			[][]int{
				{0, 1, 1},
				{1, 1, 1},
			},
		},
		&Piece{3, 2,
			[][]int{
				{0, 1},
				{1, 1},
				{1, 1},
			},
		},
		&Piece{3, 2,
			[][]int{
				{1, 0},
				{1, 1},
				{1, 1},
			},
		},
		&Piece{3, 2,
			[][]int{
				{1, 1},
				{1, 1},
				{1, 0},
			},
		},
		&Piece{3, 2,
			[][]int{
				{1, 1},
				{1, 1},
				{0, 1},
			},
		},
	},
}

/*
Leggy is great
*/
var leggy = &PieceGroup{
	[]*Piece{
		&Piece{4, 2,
			[][]int{
				{0, 1},
				{1, 1},
				{1, 0},
				{1, 0},
			},
		},
		&Piece{4, 2,
			[][]int{
				{1, 0},
				{1, 1},
				{0, 1},
				{0, 1},
			},
		},
		&Piece{4, 2,
			[][]int{
				{1, 0},
				{1, 0},
				{1, 1},
				{0, 1},
			},
		},
		&Piece{4, 2,
			[][]int{
				{0, 1},
				{0, 1},
				{1, 1},
				{1, 0},
			},
		},
		&Piece{2, 4,
			[][]int{
				{1, 1, 0, 0},
				{0, 1, 1, 1},
			},
		},
		&Piece{2, 4,
			[][]int{
				{0, 0, 1, 1},
				{1, 1, 1, 0},
			},
		},
		&Piece{2, 4,
			[][]int{
				{1, 1, 1, 0},
				{0, 0, 1, 1},
			},
		},
		&Piece{2, 4,
			[][]int{
				{0, 1, 1, 1},
				{1, 1, 0, 0},
			},
		},
	},
}

/*
Z is great
*/
var z = &PieceGroup{
	[]*Piece{
		&Piece{3, 3,
			[][]int{
				{1, 1, 0},
				{0, 1, 0},
				{0, 1, 1},
			},
		},
		&Piece{3, 3,
			[][]int{
				{0, 1, 1},
				{0, 1, 0},
				{1, 1, 0},
			},
		},
		&Piece{3, 3,
			[][]int{
				{0, 0, 1},
				{1, 1, 1},
				{1, 0, 0},
			},
		},
		&Piece{3, 3,
			[][]int{
				{1, 0, 0},
				{1, 1, 1},
				{0, 0, 1},
			},
		},
	},
}

/*
BowBow is great
*/
var bowBow = &PieceGroup{
	[]*Piece{
		&Piece{3, 3,
			[][]int{
				{1, 1, 0},
				{0, 1, 1},
				{0, 0, 1},
			},
		},
		&Piece{3, 3,
			[][]int{
				{0, 1, 1},
				{1, 1, 0},
				{1, 0, 0},
			},
		},
		&Piece{3, 3,
			[][]int{
				{0, 0, 1},
				{0, 1, 1},
				{1, 1, 0},
			},
		},
		&Piece{3, 3,
			[][]int{
				{1, 0, 0},
				{1, 1, 0},
				{0, 1, 1},
			},
		},
	},
}

/*
BigL is great
*/
var bigL = &PieceGroup{
	[]*Piece{
		&Piece{3, 3,
			[][]int{
				{1, 1, 1},
				{0, 0, 1},
				{0, 0, 1},
			},
		},
		&Piece{3, 3,
			[][]int{
				{1, 1, 1},
				{1, 0, 0},
				{1, 0, 0},
			},
		},
		&Piece{3, 3,
			[][]int{
				{0, 0, 1},
				{0, 0, 1},
				{1, 1, 1},
			},
		},
		&Piece{3, 3,
			[][]int{
				{1, 0, 0},
				{1, 0, 0},
				{1, 1, 1},
			},
		},
	},
}

/*
Plus is great
*/
var plus = &PieceGroup{
	[]*Piece{
		&Piece{3, 3,
			[][]int{
				{0, 1, 0},
				{1, 1, 1},
				{0, 1, 0},
			},
		},
	},
}

/*
L is great
*/
var l = &PieceGroup{
	[]*Piece{
		&Piece{3, 2,
			[][]int{
				{0, 1},
				{0, 1},
				{1, 1},
			},
		},
		&Piece{3, 2,
			[][]int{
				{1, 1},
				{0, 1},
				{0, 1},
			},
		},
		&Piece{3, 2,
			[][]int{
				{1, 1},
				{1, 0},
				{1, 0},
			},
		},
		&Piece{3, 2,
			[][]int{
				{1, 0},
				{1, 0},
				{1, 1},
			},
		},
		&Piece{2, 3,
			[][]int{
				{1, 0, 0},
				{1, 1, 1},
			},
		},
		&Piece{2, 3,
			[][]int{
				{0, 0, 1},
				{1, 1, 1},
			},
		},
		&Piece{2, 3,
			[][]int{
				{1, 1, 1},
				{0, 0, 1},
			},
		},
		&Piece{2, 3,
			[][]int{
				{1, 1, 1},
				{1, 0, 0},
			},
		},
	},
}

/*
Bow is great
*/
var bow = &PieceGroup{
	[]*Piece{
		&Piece{3, 2,
			[][]int{
				{0, 1},
				{1, 1},
				{1, 0},
			},
		},
		&Piece{3, 2,
			[][]int{
				{1, 0},
				{1, 1},
				{0, 1},
			},
		},
		&Piece{2, 3,
			[][]int{
				{0, 1, 1},
				{1, 1, 0},
			},
		},
		&Piece{2, 3,
			[][]int{
				{1, 1, 0},
				{0, 1, 1},
			},
		},
	},
}

/*
Block is great
*/
var block = &PieceGroup{
	[]*Piece{
		&Piece{2, 2,
			[][]int{
				{1, 1},
				{1, 1},
			},
		},
	},
}

/*
PacMan is great
*/
var pacMan = &PieceGroup{
	[]*Piece{
		&Piece{3, 2,
			[][]int{
				{1, 1},
				{0, 1},
				{1, 1},
			},
		},
		&Piece{3, 2,
			[][]int{
				{1, 1},
				{1, 0},
				{1, 1},
			},
		},
		&Piece{2, 3,
			[][]int{
				{1, 0, 1},
				{1, 1, 1},
			},
		},
		&Piece{2, 3,
			[][]int{
				{1, 1, 1},
				{1, 0, 1},
			},
		},
	},
}

/*
T is great
*/
var t = &PieceGroup{
	[]*Piece{
		&Piece{3, 3,
			[][]int{
				{1, 1, 1},
				{0, 1, 0},
				{0, 1, 0},
			},
		},
		&Piece{3, 3,
			[][]int{
				{0, 1, 0},
				{0, 1, 0},
				{1, 1, 1},
			},
		},
		&Piece{3, 3,
			[][]int{
				{0, 0, 1},
				{1, 1, 1},
				{0, 0, 1},
			},
		},
		&Piece{3, 3,
			[][]int{
				{1, 0, 0},
				{1, 1, 1},
				{1, 0, 0},
			},
		},
	},
}

/*
Oddy is great
*/
var oddy = &PieceGroup{
	[]*Piece{
		&Piece{3, 3,
			[][]int{
				{0, 1, 1},
				{1, 1, 0},
				{0, 1, 0},
			},
		},
		&Piece{3, 3,
			[][]int{
				{1, 1, 0},
				{0, 1, 1},
				{0, 1, 0},
			},
		},
		&Piece{3, 3,
			[][]int{
				{0, 1, 0},
				{0, 1, 1},
				{1, 1, 0},
			},
		},
		&Piece{3, 3,
			[][]int{
				{0, 1, 0},
				{1, 1, 0},
				{0, 1, 1},
			},
		},
		&Piece{3, 3,
			[][]int{
				{0, 0, 1},
				{1, 1, 1},
				{0, 1, 0},
			},
		},
		&Piece{3, 3,
			[][]int{
				{0, 1, 0},
				{1, 1, 1},
				{0, 0, 1},
			},
		},
		&Piece{3, 3,
			[][]int{
				{0, 1, 0},
				{1, 1, 1},
				{1, 0, 0},
			},
		},
		&Piece{3, 3,
			[][]int{
				{1, 0, 0},
				{1, 1, 1},
				{0, 1, 0},
			},
		},
	},
}

/*
Shooty is great
*/
var shooty = &PieceGroup{
	[]*Piece{
		&Piece{4, 2,
			[][]int{
				{1, 0},
				{1, 1},
				{1, 0},
				{1, 0},
			},
		},
		&Piece{4, 2,
			[][]int{
				{0, 1},
				{1, 1},
				{0, 1},
				{0, 1},
			},
		},
		&Piece{4, 2,
			[][]int{
				{1, 0},
				{1, 0},
				{1, 1},
				{1, 0},
			},
		},
		&Piece{4, 2,
			[][]int{
				{0, 1},
				{0, 1},
				{1, 1},
				{0, 1},
			},
		},
		&Piece{2, 4,
			[][]int{
				{0, 1, 0, 0},
				{1, 1, 1, 1},
			},
		},
		&Piece{2, 4,
			[][]int{
				{0, 0, 1, 0},
				{1, 1, 1, 1},
			},
		},
		&Piece{2, 4,
			[][]int{
				{1, 1, 1, 1},
				{0, 0, 1, 0},
			},
		},
		&Piece{2, 4,
			[][]int{
				{1, 1, 1, 1},
				{0, 1, 0, 0},
			},
		},
	},
}

/*
LongL is great
*/
var longL = &PieceGroup{
	[]*Piece{
		&Piece{4, 2,
			[][]int{
				{0, 1},
				{0, 1},
				{0, 1},
				{1, 1},
			},
		},
		&Piece{4, 2,
			[][]int{
				{1, 1},
				{0, 1},
				{0, 1},
				{0, 1},
			},
		},
		&Piece{4, 2,
			[][]int{
				{1, 1},
				{1, 0},
				{1, 0},
				{1, 0},
			},
		},
		&Piece{4, 2,
			[][]int{
				{1, 0},
				{1, 0},
				{1, 0},
				{1, 1},
			},
		},
		&Piece{2, 4,
			[][]int{
				{1, 0, 0, 0},
				{1, 1, 1, 1},
			},
		},
		&Piece{2, 4,
			[][]int{
				{0, 0, 0, 1},
				{1, 1, 1, 1},
			},
		},
		&Piece{2, 4,
			[][]int{
				{1, 1, 1, 1},
				{0, 0, 0, 1},
			},
		},
		&Piece{2, 4,
			[][]int{
				{1, 1, 1, 1},
				{1, 0, 0, 0},
			},
		},
	},
}

/*
Tetris is great
*/
var tetris = &PieceGroup{
	[]*Piece{
		&Piece{2, 3,
			[][]int{
				{1, 1, 1},
				{0, 1, 0},
			},
		},
		&Piece{2, 3,
			[][]int{
				{0, 1, 0},
				{1, 1, 1},
			},
		},
		&Piece{3, 2,
			[][]int{
				{0, 1},
				{1, 1},
				{0, 1},
			},
		},
		&Piece{3, 2,
			[][]int{
				{1, 0},
				{1, 1},
				{1, 0},
			},
		},
	},
}
