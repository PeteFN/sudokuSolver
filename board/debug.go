package board

import "fmt"

type boardPos struct {
	row int
	col int
}

var debugOn = false

var cellMap = make(map[*Cell]*boardPos)

func SetDebug(b *Board) {
	if debugOn {
		for i := 0; i < b.GetNumRows(); i++ {
			for j := 0; j < b.GetNumCols(); j++ {
				bp := &boardPos{
					row: i,
					col: j,
				}
				cellMap[b.cells[i][j]] = bp
			}
		}
	}
}

func logConstraint(c *Cell, lc *Cell, val int) {
	if debugOn {
		if cPos, found := cellMap[c]; found {
			if constraintPos, found := cellMap[lc]; found {
				fmt.Printf("Setting value at %d,%d to %d, constraing cell at %d,%d\n", cPos.row, cPos.col, val, constraintPos.row, constraintPos.col)
			} else {
				fmt.Printf("INVALID: NO DATA FOR DEBUG\n")
			}
		} else {
			fmt.Printf("INVALID: NO DATA FOR DEBUG\n")
		}
	}
}

func logUnconstraint(c *Cell, lc *Cell, val int) {
	if debugOn {
		if cPos, found := cellMap[c]; found {
			if constraintPos, found := cellMap[lc]; found {
				fmt.Printf("Unsetting value at %d,%d to %d, removing constraing cell at %d,%d\n", cPos.row, cPos.col, val, constraintPos.row, constraintPos.col)
			} else {
				fmt.Printf("INVALID: NO DATA FOR DEBUG\n")
			}
		} else {
			fmt.Printf("INVALID: NO DATA FOR DEBUG\n")
		}
	}
}
