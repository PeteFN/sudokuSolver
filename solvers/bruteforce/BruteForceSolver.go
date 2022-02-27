package bruteforce

import (
	"fmt"
	"sudokuSolver/board"
)

var _logAttempts = false

type Solver struct {
	board              *board.Board
	cellTrackers       []*CellTracker
	currentTrackersPos int
	currentSolution    []*ValueAttempt
	currentCellPos     int
	currentSolutionPos int

	numValueAttempts      int
	numValidValueAttempts int
	numRollbacks          int
}

func NewSolver(board *board.Board) *Solver {
	cellTrackers := make([]*CellTracker, board.GetNumUnknown())
	solver := &Solver{
		board:              board,
		cellTrackers:       cellTrackers,
		currentTrackersPos: -1, //Set to -1 as nextTracker will increment it
		currentCellPos:     -1,
		currentSolution:    make([]*ValueAttempt, board.GetNumUnknown()),
		currentSolutionPos: -1,
	}
	solver._nextTracker()
	return solver
}

func (s *Solver) Solve() bool {
	tracker := s.cellTrackers[s.currentTrackersPos]
	va := tracker.nextAttempt()
	if va != nil {
		isSuccessful := s._attemptValue(va)
		if isSuccessful {
			s.logAttempt()
			if s.board.IsSolved() {
				return true
			}
			s._nextTracker()
		}
		s.Solve()
	} else {
		s._undoLastAttempt()
		s._removeTracker()
		s.Solve()
	}
	return false
}

func (s *Solver) _nextTracker() {
	s.currentCellPos++
	row, col := s._resolveRowAndCol()
	for s.board.GetValue(row, col) != 0 {
		s.currentCellPos++
		row, col = s._resolveRowAndCol()
	}
	curCellPossibleValues := s.board.GetPossibleValues(row, col)
	tracker := newCellTracker(row, col, curCellPossibleValues)
	s._addTracker(tracker)
}

func (s *Solver) _resolveRowAndCol() (int, int) {
	row := s.currentCellPos / s.board.GetNumRows()
	col := s.currentCellPos % s.board.GetNumRows()
	return row, col
}

func (s *Solver) _removeTracker() {
	s.currentTrackersPos--
}

func (s *Solver) _addTracker(tracker *CellTracker) {
	s.currentTrackersPos++
	s.cellTrackers[s.currentTrackersPos] = tracker
}

func (s *Solver) _attemptValue(attempt *ValueAttempt) bool {
	s.numValueAttempts++
	attemptSuccessful := s.board.AttemptValue(attempt.row, attempt.col, attempt.val)
	if attemptSuccessful {
		s.numValidValueAttempts++
		s.currentSolutionPos++
		s.currentSolution[s.currentSolutionPos] = attempt
	}
	return attemptSuccessful
}

func (s *Solver) _undoLastAttempt() {
	s.numRollbacks++
	lastAttempt := s.currentSolution[s.currentSolutionPos]
	s.currentSolutionPos--
	s.currentCellPos = lastAttempt.row*s.board.GetNumRows() + lastAttempt.col
	s.board.UndoAttempt(lastAttempt.row, lastAttempt.col)
}

func (s *Solver) logAttempt() {
	if _logAttempts {
		lastAttempt := s.currentSolution[s.currentSolutionPos]
		fmt.Printf("=========================\n")
		fmt.Printf("=======Last Attempt======\n")
		fmt.Printf("  row=%d, col=%d, val=%d\n", lastAttempt.row, lastAttempt.col, lastAttempt.val)
		fmt.Printf("=========================\n")
		fmt.Printf("=========[BOARD]=========\n")
		fmt.Printf("%s", s.board)
		fmt.Printf("=========================\n")
	}
}

func (s *Solver) LogSolution() {
	fmt.Printf("==========================\n")
	fmt.Printf("=========[SOLVED]=========\n")
	fmt.Printf("=assingedValueAttempts=%d=\n", s.numValueAttempts)
	fmt.Printf("=validAssignedValues=%d===\n", s.numValueAttempts)
	fmt.Printf("=numRollbacks=%d==========\n", s.numRollbacks)
	fmt.Printf("==========================\n")
	fmt.Printf("=========[BOARD]==========\n")
	fmt.Printf("%s", s.board)
	fmt.Printf("=========================\n")
}

func (s *Solver) ValidateSolution() bool {
	return s.board.ValidateSolution()
}
