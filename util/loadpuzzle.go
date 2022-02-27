package util

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"sudokuSolver/board"
)

const (
	informationMode = "[INFORMATION]"
	dimensionsLabel = "DIMENSIONS"
	numGroupsLabel  = "NUM_GROUPS"
	puzzleMode      = "[PUZZLE]"
	groupMode       = "[GROUPS]"
	none            = "NONE"
)

func LoadPuzzle(fileName string) *board.Board {
	f, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	mode := none

	var vals [][]int
	var groups [][]int
	var numGroups int
	var numRows int
	var numCols int
	var curRow int
	groupCnt := 0
	for scanner.Scan() {
		line := scanner.Text()
		if line == informationMode {
			mode = informationMode
		} else if line == puzzleMode {
			vals = make([][]int, numRows)
			for i := 0; i < numRows; i++ {
				vals[i] = make([]int, numCols)
			}
			curRow = 0
			mode = puzzleMode
		} else if line == groupMode {
			groups = make([][]int, numRows)
			for i := 0; i < numRows; i++ {
				groups[i] = make([]int, numCols)
			}
			curRow = 0
			mode = groupMode
		} else {
			if mode == informationMode {
				if strings.HasPrefix(line, dimensionsLabel) {
					valStr := strings.Split(line, "=")[1]
					dims := strings.Split(valStr, "x")
					numRows, err = strconv.Atoi(dims[0])
					if err != nil {
						panic(err)
					}
					numCols, err = strconv.Atoi(dims[1])
					if err != nil {
						panic(err)
					}
				} else if strings.HasPrefix(line, numGroupsLabel) {
					valStr := strings.Split(line, "=")[1]
					numGroups, err = strconv.Atoi(valStr)
					if err != nil {
						panic(err)
					}
				}
			} else if mode == puzzleMode {
				line = scanner.Text()
				for i, cInt := range line {
					val, err := strconv.Atoi(string(cInt))
					if err != nil {
						panic(err)
					}
					vals[curRow][i] = val
				}
				curRow++
			} else if mode == groupMode {
				line = scanner.Text()
				for i, cInt := range line {
					val, err := strconv.Atoi(string(cInt))
					if err != nil {
						panic(err)
					}
					groups[curRow][i] = val
					groupCnt = MathMaxInt(groupCnt, val)
				}
				curRow++
			}
		}
	}

	return board.NewBoard(vals, groups, numGroups)
}
