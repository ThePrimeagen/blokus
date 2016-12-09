package blokus

import (
	"fmt"
	"math"
	"strconv"
)

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
	isOne bool
	isTwo bool
	Value [][]int
}

func (p *Piece) String() string {
	outStr := ""

	for row := 0; row < p.Row; row++ {
		for col := 0; col < p.Col; col++ {
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
	Row         int
	Col         int
	count       int
	solvedCount int
	ones        int
	twos        int
	hasOne      bool
	hasTwo      bool
	Value       [][]int
}

/*
Solved is great
*/
func (b *Board) Solved() bool {
	return b.count == b.solvedCount
}

func (b *Board) String() string {
	outStr := fmt.Sprintf("Board(%v) {\n", b.ones)

	for row := 0; row < b.Row; row++ {
		outStr += "{"
		for col := 0; col < b.Col; col++ {
			outStr += strconv.Itoa(b.Value[row][col])
			if col < b.Col-1 {
				outStr += ","
			}
		}

		outStr += "},\n"
	}

	outStr += "}\n"
	return outStr
}

/*
NewBoard is great
*/
func NewBoard(size int) *Board {
	board := Board{size, size, 0, 0, 0, 0, false, false, [][]int{}}
	rows := make([][]int, size)

	boardValue := rows[:]

	for i := range rows {
		row := make([]int, size)
		if i == 0 || i == size-1 {
			for j := 0; j < size; j++ {
				row[j] = 1
			}
		}

		// the corners should be at 0.2
		row[0]++
		row[size-1]++

		rows[i] = row[0:]
	}

	board.Value = boardValue

	board.solvedCount = size * size
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

	if p.isOne {
		b.hasOne = true
	}

	if p.isTwo {
		b.hasTwo = true
	}

	for r := 0; r < p.Row; r++ {
		for c := 0; c < p.Col; c++ {
			if p.Value[r][c] == 0 {
				continue
			}

			offsetR := row + r
			offsetC := col + c

			b.count++
			b.Value[offsetR][offsetC] += 10

			onesAndTwos(b, offsetR, offsetC)
		}
	}

	if b.count > b.solvedCount {
		fmt.Printf("SUPER ERROR \n%v", p)
		b.Value = b.Value[100:]
	}
	return true
}

/*
Remove same piece from same position
*/
func (b *Board) Remove(p *Piece, row, col int) bool {

	if p.isOne {
		b.hasOne = false
	}

	if p.isTwo {
		b.hasTwo = false
	}

	for r := 0; r < p.Row; r++ {
		for c := 0; c < p.Col; c++ {
			if p.Value[r][c] == 0 {
				continue
			}

			offsetR := row + r
			offsetC := col + c

			b.count--
			b.Value[offsetR][offsetC] -= 10

			decOnesAndTwos(b, offsetR, offsetC)
		}
	}

	return true
}

func decOnesAndTwos(b *Board, row, col int) {
	if b.Value[row][col] == 4 {
		b.ones++
	}
	if b.Value[row][col] == 3 {
		incDecTwos(b, row, col, 1)
	}

	if row-1 >= 0 {
		val := b.Value[row-1][col] - 1
		if val == 3 {
			b.ones--
			incDecTwos(b, row-1, col, 1)
		}
		b.Value[row-1][col] = val
	}
	if row+1 < b.Row {
		val := b.Value[row+1][col] - 1
		if val == 3 {
			b.ones--
			incDecTwos(b, row+1, col, 1)
		}
		b.Value[row+1][col] = val
	}
	if col-1 >= 0 {
		val := b.Value[row][col-1] - 1
		if val == 3 {
			b.ones--
			incDecTwos(b, row, col-1, 1)
		}
		b.Value[row][col-1] = val
	}
	if col+1 < b.Col {
		val := b.Value[row][col+1] - 1
		if val == 3 {
			b.ones--
			incDecTwos(b, row, col+1, 1)
		}
		b.Value[row][col+1] = val
	}
}

func onesAndTwos(b *Board, row, col int) {
	if b.Value[row][col] == 13 {
		incDecTwos(b, row, col, -1)
	}
	if b.Value[row][col] == 14 {
		b.ones--
	}

	if row-1 >= 0 {
		val := b.Value[row-1][col] + 1
		if val == 4 {
			b.ones++
			incDecTwos(b, row-1, col, -1)
		}
		if val == 3 {
			incDecTwos(b, row-1, col, 1)
		}
		b.Value[row-1][col] = val
	}
	if row+1 < b.Row {
		val := b.Value[row+1][col] + 1
		if val == 4 {
			b.ones++
			incDecTwos(b, row+1, col, -1)
		}
		if val == 3 {
			incDecTwos(b, row+1, col, 1)
		}
		b.Value[row+1][col] = val
	}
	if col-1 >= 0 {
		val := b.Value[row][col-1] + 1
		if val == 4 {
			b.ones++
			incDecTwos(b, row, col-1, -1)
		}
		if val == 3 {
			incDecTwos(b, row, col-1, 1)
		}
		b.Value[row][col-1] = val
	}
	if col+1 < b.Col {
		val := b.Value[row][col+1] + 1
		if val == 4 {
			b.ones++
			incDecTwos(b, row, col+1, -1)
		}
		if val == 3 {
			incDecTwos(b, row, col+1, 1)
		}
		b.Value[row][col+1] = val
	}
}

// takes a 3 value square and checks its adjancents for another 3
func incDecTwos(b *Board, row, col, dir int) {
	if row-1 >= 0 {
		val := b.Value[row-1][col]
		if val == 3 {
			b.twos = b.twos + dir
			return
		}
	}
	if row+1 < b.Row {
		val := b.Value[row+1][col]
		if val == 3 {
			b.twos = b.twos + dir
			return
		}
	}
	if col-1 >= 0 {
		val := b.Value[row][col-1]
		if val == 3 {
			b.twos = b.twos + dir
			return
		}
	}
	if col+1 < b.Col {
		val := b.Value[row][col+1]
		if val == 3 {
			b.twos = b.twos + dir
			return
		}
	}
}

/*
IsSolvable thats more like it
*/
func (b *Board) IsSolvable() bool {
	if b.hasOne && b.ones > 0 {
		return false
	}

	if b.hasTwo && b.twos > 0 {
		return false
	}

	return b.ones < 2 && b.twos < 2
}

/*
Key thats more like it
*/
func (b *Board) Key() string {
	cKey := 0
	rKey := 0

	// go 64 bit system
	for r := 0; r < b.Row; r++ {
		for c := 0; c < b.Col; c++ {
			if b.Value[r][c] == 1 {
				cKey += int(math.Pow10(c + 1))
				rKey += int(math.Pow10(r + 1))
			}
		}
	}

	return string(cKey) + ":" + string(rKey)
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

	valid := true
	for r := 0; r < p.Row && valid; r++ {
		for c := 0; c < p.Col && valid; c++ {
			if p.Value[r][c] == 0 {
				continue
			}

			offsetR := row + r
			offsetC := col + c

			valid = b.Value[offsetR][offsetC] < 5
		}
	}

	return valid

}

/*
OneByOne is great
*/
var oneByOne = &PieceGroup{
	[]*Piece{
		&Piece{1, 1,
			true,
			false,
			[][]int{{1}},
		},
	},
}

/*
OneByTwo is great
*/
var oneByTwo = &PieceGroup{
	[]*Piece{
		&Piece{2, 1,
			false,
			true,
			[][]int{
				{1},
				{1},
			},
		},
		&Piece{1, 2,
			false,
			false,
			[][]int{{1, 1}},
		},
	},
}

/*
OneByThree is great
*/
var oneByThree = &PieceGroup{
	[]*Piece{
		&Piece{3, 1,
			false,
			false,
			[][]int{
				{1},
				{1},
				{1},
			},
		},
		&Piece{1, 3,
			false,
			false,
			[][]int{{1, 1, 1}},
		},
	},
}

/*
OneByFour is great
*/
var oneByFour = &PieceGroup{
	[]*Piece{
		&Piece{4, 1,
			false,
			false,
			[][]int{
				{1},
				{1},
				{1},
				{1},
			},
		},
		&Piece{1, 4,
			false,
			false,
			[][]int{{1, 1, 1, 1}},
		},
	},
}

/*
OneByFive is great
*/
var oneByFive = &PieceGroup{
	[]*Piece{
		&Piece{5, 1,
			false,
			false,
			[][]int{
				{1},
				{1},
				{1},
				{1},
				{1},
			},
		},
		&Piece{1, 5,
			false,
			false,
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
			false,
			false,
			[][]int{
				{1, 0},
				{1, 1},
			},
		},
		&Piece{2, 2,
			false,
			false,
			[][]int{
				{1, 1},
				{0, 1},
			},
		},
		&Piece{2, 2,
			false,
			false,
			[][]int{
				{0, 1},
				{1, 1},
			},
		},
		&Piece{2, 2,
			false,
			false,
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
			false,
			false,
			[][]int{
				{1, 1, 0},
				{1, 1, 1},
			},
		},
		&Piece{2, 3,
			false,
			false,
			[][]int{
				{1, 1, 1},
				{1, 1, 0},
			},
		},
		&Piece{2, 3,
			false,
			false,
			[][]int{
				{1, 1, 1},
				{0, 1, 1},
			},
		},
		&Piece{2, 3,
			false,
			false,
			[][]int{
				{0, 1, 1},
				{1, 1, 1},
			},
		},
		&Piece{3, 2,
			false,
			false,
			[][]int{
				{0, 1},
				{1, 1},
				{1, 1},
			},
		},
		&Piece{3, 2,
			false,
			false,
			[][]int{
				{1, 0},
				{1, 1},
				{1, 1},
			},
		},
		&Piece{3, 2,
			false,
			false,
			[][]int{
				{1, 1},
				{1, 1},
				{1, 0},
			},
		},
		&Piece{3, 2,
			false,
			false,
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
			false,
			false,
			[][]int{
				{0, 1},
				{1, 1},
				{1, 0},
				{1, 0},
			},
		},
		&Piece{4, 2,
			false,
			false,
			[][]int{
				{1, 0},
				{1, 1},
				{0, 1},
				{0, 1},
			},
		},
		&Piece{4, 2,
			false,
			false,
			[][]int{
				{1, 0},
				{1, 0},
				{1, 1},
				{0, 1},
			},
		},
		&Piece{4, 2,
			false,
			false,
			[][]int{
				{0, 1},
				{0, 1},
				{1, 1},
				{1, 0},
			},
		},
		&Piece{2, 4,
			false,
			false,
			[][]int{
				{1, 1, 0, 0},
				{0, 1, 1, 1},
			},
		},
		&Piece{2, 4,
			false,
			false,
			[][]int{
				{0, 0, 1, 1},
				{1, 1, 1, 0},
			},
		},
		&Piece{2, 4,
			false,
			false,
			[][]int{
				{1, 1, 1, 0},
				{0, 0, 1, 1},
			},
		},
		&Piece{2, 4,
			false,
			false,
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
			false,
			false,
			[][]int{
				{1, 1, 0},
				{0, 1, 0},
				{0, 1, 1},
			},
		},
		&Piece{3, 3,
			false,
			false,
			[][]int{
				{0, 1, 1},
				{0, 1, 0},
				{1, 1, 0},
			},
		},
		&Piece{3, 3,
			false,
			false,
			[][]int{
				{0, 0, 1},
				{1, 1, 1},
				{1, 0, 0},
			},
		},
		&Piece{3, 3,
			false,
			false,
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
			false,
			false,
			[][]int{
				{1, 1, 0},
				{0, 1, 1},
				{0, 0, 1},
			},
		},
		&Piece{3, 3,
			false,
			false,
			[][]int{
				{0, 1, 1},
				{1, 1, 0},
				{1, 0, 0},
			},
		},
		&Piece{3, 3,
			false,
			false,
			[][]int{
				{0, 0, 1},
				{0, 1, 1},
				{1, 1, 0},
			},
		},
		&Piece{3, 3,
			false,
			false,
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
			false,
			false,
			[][]int{
				{1, 1, 1},
				{0, 0, 1},
				{0, 0, 1},
			},
		},
		&Piece{3, 3,
			false,
			false,
			[][]int{
				{1, 1, 1},
				{1, 0, 0},
				{1, 0, 0},
			},
		},
		&Piece{3, 3,
			false,
			false,
			[][]int{
				{0, 0, 1},
				{0, 0, 1},
				{1, 1, 1},
			},
		},
		&Piece{3, 3,
			false,
			false,
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
			false,
			false,
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
			false,
			false,
			[][]int{
				{0, 1},
				{0, 1},
				{1, 1},
			},
		},
		&Piece{3, 2,
			false,
			false,
			[][]int{
				{1, 1},
				{0, 1},
				{0, 1},
			},
		},
		&Piece{3, 2,
			false,
			false,
			[][]int{
				{1, 1},
				{1, 0},
				{1, 0},
			},
		},
		&Piece{3, 2,
			false,
			false,
			[][]int{
				{1, 0},
				{1, 0},
				{1, 1},
			},
		},
		&Piece{2, 3,
			false,
			false,
			[][]int{
				{1, 0, 0},
				{1, 1, 1},
			},
		},
		&Piece{2, 3,
			false,
			false,
			[][]int{
				{0, 0, 1},
				{1, 1, 1},
			},
		},
		&Piece{2, 3,
			false,
			false,
			[][]int{
				{1, 1, 1},
				{0, 0, 1},
			},
		},
		&Piece{2, 3,
			false,
			false,
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
			false,
			false,
			[][]int{
				{0, 1},
				{1, 1},
				{1, 0},
			},
		},
		&Piece{3, 2,
			false,
			false,
			[][]int{
				{1, 0},
				{1, 1},
				{0, 1},
			},
		},
		&Piece{2, 3,
			false,
			false,
			[][]int{
				{0, 1, 1},
				{1, 1, 0},
			},
		},
		&Piece{2, 3,
			false,
			false,
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
			false,
			false,
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
			false,
			false,
			[][]int{
				{1, 1},
				{0, 1},
				{1, 1},
			},
		},
		&Piece{3, 2,
			false,
			false,
			[][]int{
				{1, 1},
				{1, 0},
				{1, 1},
			},
		},
		&Piece{2, 3,
			false,
			false,
			[][]int{
				{1, 0, 1},
				{1, 1, 1},
			},
		},
		&Piece{2, 3,
			false,
			false,
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
			false,
			false,
			[][]int{
				{1, 1, 1},
				{0, 1, 0},
				{0, 1, 0},
			},
		},
		&Piece{3, 3,
			false,
			false,
			[][]int{
				{0, 1, 0},
				{0, 1, 0},
				{1, 1, 1},
			},
		},
		&Piece{3, 3,
			false,
			false,
			[][]int{
				{0, 0, 1},
				{1, 1, 1},
				{0, 0, 1},
			},
		},
		&Piece{3, 3,
			false,
			false,
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
			false,
			false,
			[][]int{
				{0, 1, 1},
				{1, 1, 0},
				{0, 1, 0},
			},
		},
		&Piece{3, 3,
			false,
			false,
			[][]int{
				{1, 1, 0},
				{0, 1, 1},
				{0, 1, 0},
			},
		},
		&Piece{3, 3,
			false,
			false,
			[][]int{
				{0, 1, 0},
				{0, 1, 1},
				{1, 1, 0},
			},
		},
		&Piece{3, 3,
			false,
			false,
			[][]int{
				{0, 1, 0},
				{1, 1, 0},
				{0, 1, 1},
			},
		},
		&Piece{3, 3,
			false,
			false,
			[][]int{
				{0, 0, 1},
				{1, 1, 1},
				{0, 1, 0},
			},
		},
		&Piece{3, 3,
			false,
			false,
			[][]int{
				{0, 1, 0},
				{1, 1, 1},
				{0, 0, 1},
			},
		},
		&Piece{3, 3,
			false,
			false,
			[][]int{
				{0, 1, 0},
				{1, 1, 1},
				{1, 0, 0},
			},
		},
		&Piece{3, 3,
			false,
			false,
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
			false,
			false,
			[][]int{
				{1, 0},
				{1, 1},
				{1, 0},
				{1, 0},
			},
		},
		&Piece{4, 2,
			false,
			false,
			[][]int{
				{0, 1},
				{1, 1},
				{0, 1},
				{0, 1},
			},
		},
		&Piece{4, 2,
			false,
			false,
			[][]int{
				{1, 0},
				{1, 0},
				{1, 1},
				{1, 0},
			},
		},
		&Piece{4, 2,
			false,
			false,
			[][]int{
				{0, 1},
				{0, 1},
				{1, 1},
				{0, 1},
			},
		},
		&Piece{2, 4,
			false,
			false,
			[][]int{
				{0, 1, 0, 0},
				{1, 1, 1, 1},
			},
		},
		&Piece{2, 4,
			false,
			false,
			[][]int{
				{0, 0, 1, 0},
				{1, 1, 1, 1},
			},
		},
		&Piece{2, 4,
			false,
			false,
			[][]int{
				{1, 1, 1, 1},
				{0, 0, 1, 0},
			},
		},
		&Piece{2, 4,
			false,
			false,
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
			false,
			false,
			[][]int{
				{0, 1},
				{0, 1},
				{0, 1},
				{1, 1},
			},
		},
		&Piece{4, 2,
			false,
			false,
			[][]int{
				{1, 1},
				{0, 1},
				{0, 1},
				{0, 1},
			},
		},
		&Piece{4, 2,
			false,
			false,
			[][]int{
				{1, 1},
				{1, 0},
				{1, 0},
				{1, 0},
			},
		},
		&Piece{4, 2,
			false,
			false,
			[][]int{
				{1, 0},
				{1, 0},
				{1, 0},
				{1, 1},
			},
		},
		&Piece{2, 4,
			false,
			false,
			[][]int{
				{1, 0, 0, 0},
				{1, 1, 1, 1},
			},
		},
		&Piece{2, 4,
			false,
			false,
			[][]int{
				{0, 0, 0, 1},
				{1, 1, 1, 1},
			},
		},
		&Piece{2, 4,
			false,
			false,
			[][]int{
				{1, 1, 1, 1},
				{0, 0, 0, 1},
			},
		},
		&Piece{2, 4,
			false,
			false,
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
			false,
			false,
			[][]int{
				{1, 1, 1},
				{0, 1, 0},
			},
		},
		&Piece{2, 3,
			false,
			false,
			[][]int{
				{0, 1, 0},
				{1, 1, 1},
			},
		},
		&Piece{3, 2,
			false,
			false,
			[][]int{
				{0, 1},
				{1, 1},
				{0, 1},
			},
		},
		&Piece{3, 2,
			false,
			false,
			[][]int{
				{1, 0},
				{1, 1},
				{1, 0},
			},
		},
	},
}
