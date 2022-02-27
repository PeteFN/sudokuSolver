package bruteforce

type CellTracker struct {
	attempts map[*ValueAttempt]bool
}

func newCellTracker(row int, col int, vals []int) *CellTracker {
	attempts := make(map[*ValueAttempt]bool)
	for _, val := range vals {
		attempts[&ValueAttempt{
			row: row,
			col: col,
			val: val,
		}] = false
	}
	return &CellTracker{
		attempts: attempts,
	}
}

func (ct *CellTracker) nextAttempt() *ValueAttempt {
	for k, _ := range ct.attempts {
		delete(ct.attempts, k)
		return k
	}
	return nil
}
