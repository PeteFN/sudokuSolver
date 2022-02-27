package board

import (
	"fmt"
	"strings"
)

type Board struct {
	groups           map[int]*CellList
	cells            [][]*Cell
	numUnknownValues int
}

func (b *Board) GetNumUnknown() int {
	return b.numUnknownValues
}

func (b *Board) IsSolved() bool {
	return b.numUnknownValues == 0
}

func (b *Board) String() string {
	var sb strings.Builder
	for i := 0; i < len(b.cells); i++ {
		for j := 0; j < len(b.cells[i]); j++ {
			fmt.Fprintf(&sb, "%d", b.cells[i][j].val)
		}
		fmt.Fprintf(&sb, "\n")
	}
	return sb.String()
}

func NewBoard(vals [][]int, groups [][]int, numGroups int) *Board {
	numRows := len(vals)
	numCols := len(vals[0])

	groupMapping := make(map[int]*CellList)
	groupByCell := make(map[*Cell]int)
	for i := 1; i <= numGroups; i++ {
		groupMapping[i] = newCellList()
	}

	colMapping := make(map[int]*CellList)
	for i := 0; i < numCols; i++ {
		colMapping[i] = newCellList()
	}

	cells := make([][]*Cell, numRows)
	for i := 0; i < numRows; i++ {
		row := make([]*Cell, numCols)
		for j, _ := range row {
			c := newCell()
			row[j] = c
			groupMapping[groups[i][j]].add(c) //add to group
			colMapping[j].add(c)              //add to columns
			groupByCell[c] = groups[i][j]     //Associate the cell to the group number
		}
		cells[i] = row
	}

	//Link rows
	for i := 0; i < numRows; i++ {
		for j := 0; j < numCols; j++ {
			for z := j + 1; z < numCols; z++ {
				if cells[i][j] == cells[i][z] {
					continue
				}
				cells[i][j].linkCell(cells[i][z])
				cells[i][z].linkCell(cells[i][j])
			}
		}
	}

	//Link cols
	for i := 0; i < numRows; i++ {
		for j := 0; j < numCols; j++ {
			for z := j + 1; z < numCols; z++ {
				if cells[j][i] == cells[z][i] {
					continue
				}
				cells[j][i].linkCell(cells[z][i])
				cells[z][i].linkCell(cells[j][i])
			}
		}
	}

	//Link groups
	for i := 0; i < numRows; i++ {
		for j := 0; j < numCols; j++ {
			for _, lc := range groupMapping[groupByCell[cells[i][j]]].iterate() {
				if cells[i][j] == lc {
					continue
				}
				cells[i][j].linkCell(lc)
				lc.linkCell(cells[i][j])
			}
		}
	}

	//Set initial constraints and get number to solve
	numToSolve := 0
	for i := 0; i < numRows; i++ {
		for j := 0; j < numCols; j++ {
			if vals[i][j] != 0 {
				cells[i][j].assignVal(vals[i][j])
			} else {
				numToSolve++
			}
		}
	}

	return &Board{
		cells:            cells,
		numUnknownValues: numToSolve,
		groups:           groupMapping,
	}
}

func (b *Board) AttemptValue(row int, col int, val int) bool {
	isSuccessful := b.cells[row][col].assignVal(val)
	if isSuccessful {
		b.numUnknownValues--
	}
	return isSuccessful
}

func (b *Board) UndoAttempt(row int, col int) {
	b.numUnknownValues++
	b.cells[row][col].unassignVal()
}

func (b *Board) GetValue(row int, col int) int {
	return b.cells[row][col].val
}

func (b *Board) GetNumRows() int {
	return len(b.cells)
}

func (b *Board) GetNumCols() int {
	return len(b.cells[0])
}

func (b *Board) GetPossibleValues(row int, col int) []int {
	return b.cells[row][col].getPossibleValues()
}

func (b *Board) ValidateSolution() bool {

	//validate rows
	for i := 0; i < len(b.cells); i++ {
		validator := new_validator(len(b.cells))
		for j := 0; j < len(b.cells[i]); j++ {
			validator.addVal(b.cells[i][j].val)
		}
		if !validator.validate() {
			fmt.Printf("Failed validation due to rows at row: %d\n", i)
			return false
		}
	}

	//validate cols
	for i := 0; i < len(b.cells[0]); i++ {
		validator := new_validator(len(b.cells[0]))
		for j := 0; j < len(b.cells); j++ {
			validator.addVal(b.cells[j][i].val)
		}
		if !validator.validate() {
			fmt.Printf("Failed validation due to cols at col: %d\n", i)
			return false
		}
	}

	//validate groups
	for groupNum, group := range b.groups {
		validator := new_validator(group.size())
		for _, c := range group.iterate() {
			validator.addVal(c.val)
		}
		if !validator.validate() {
			fmt.Printf("Failed validation due to groups, group: %d\n", groupNum)
			return false
		}
	}

	return true
}

type _validator struct {
	vals map[int]bool
}

func new_validator(size int) *_validator {
	validator := &_validator{
		vals: make(map[int]bool),
	}
	for i := 1; i < size+1; i++ {
		validator.vals[i] = false
	}
	return validator
}

func (v *_validator) addVal(val int) {
	v.vals[val] = true
}

func (v *_validator) validate() bool {
	for _, val := range v.vals {
		if !val {
			return false
		}
	}
	return true
}
