package board

type CellConstraints struct {
	constraintsByCell  map[*Cell]int
	constraintsByValue map[int]*CellSet
}

func newCellConstraints() *CellConstraints {
	cs := &CellConstraints{
		constraintsByCell:  make(map[*Cell]int),
		constraintsByValue: make(map[int]*CellSet),
	}
	for i := 1; i <= 9; i++ {
		cs.constraintsByValue[i] = newCellSet()
	}
	return cs
}

func (cc *CellConstraints) possibleValues() []int {
	baseArray := make([]int, 9)
	idx := 0
	for i := 1; i <= 9; i++ {
		//fmt.Printf("%d\n", i)
		if cc.constraintsByValue[i].size() == 0 {
			baseArray[idx] = i
			idx++
		}
	}
	return baseArray[:idx]
}

func (cc *CellConstraints) isPossibleValue(val int) bool {
	return cc.constraintsByValue[val].size() == 0
}

func (cc *CellConstraints) addConstraint(val int, cell *Cell) {
	cc.constraintsByCell[cell] = val
	cc.constraintsByValue[val].add(cell)
}

func (cc *CellConstraints) removeConstraintByCell(cell *Cell) {
	val := cc.constraintsByCell[cell]
	delete(cc.constraintsByCell, cell)
	cc.constraintsByValue[val].remove(cell)
}
