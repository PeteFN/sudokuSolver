package board

type CellSet struct {
	set map[*Cell]bool
}

func newCellSet() *CellSet {
	return &CellSet{
		set: make(map[*Cell]bool),
	}
}

func (cs *CellSet) add(c *Cell) bool {
	found, _ := cs.set[c]
	cs.set[c] = true
	return found
}

func (cs *CellSet) remove(c *Cell) bool {
	found, _ := cs.set[c]
	delete(cs.set, c)
	return found
}

func (cs *CellSet) size() int {
	return len(cs.set)
}

func (cs *CellSet) iterate() chan *Cell {
	iterable := make(chan *Cell)
	go func() {
		for k, _ := range cs.set {
			iterable <- k
		}
		close(iterable)
	}()
	return iterable
}
