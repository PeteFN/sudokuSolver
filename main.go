package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	board2 "sudokuSolver/board"
	"sudokuSolver/solvers/bruteforce"
	"sudokuSolver/util"
)

var puzzle = "./puzzles/easy1.puzzle"

func main() {
	pFile := flag.String("pFile", "", "If given, will only solve the puzzle at the given file")
	flag.Parse()
	if *pFile != "" {
		solveSingle(*pFile)
	} else {
		solveAll()
	}
}

func solveSingle(fileName string) {
	solvePuzzle(fileName)
}

func solveAll() {
	filepath.Walk("./puzzles", func(path string, info os.FileInfo, err error) error {
		if !info.Mode().IsDir() && filepath.Ext(info.Name()) == ".puzzle" {
			fmt.Printf("Solving puzzle at %s\n", path)
			solvePuzzle(path)
		}
		return nil
	})
	fmt.Printf("DONE")
}

func solvePuzzle(puzzleFile string) {
	board := util.LoadPuzzle(puzzleFile)
	board2.SetDebug(board)
	solver := bruteforce.NewSolver(board)
	isSolved := solver.Solve()
	if isSolved {
		solver.LogSolution()
	} else {
		panic("UNSOLVABLE PUZZLE")
		fmt.Printf("INVALID_SOLUTION")
	}
}
