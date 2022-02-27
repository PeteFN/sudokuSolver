package board

type Cell struct {
	//row int
	//col int
	val         int
	LinkedCells *CellSet
	Constraints *CellConstraints
}

func newCell() *Cell {
	return &Cell{
		val:         0,
		LinkedCells: newCellSet(),
		Constraints: newCellConstraints(),
	}
}

func (c *Cell) linkCell(lc *Cell) {
	c.LinkedCells.add(lc)
}

func (c *Cell) getPossibleValues() []int {
	return c.Constraints.possibleValues()
}

func (c *Cell) assignVal(val int) bool {
	if c.Constraints.isPossibleValue(val) {
		c.val = val
		for lc := range c.LinkedCells.iterate() {
			logConstraint(c, lc, val)
			lc.constrain(val, c)
		}
		return true
	}
	return false
}

func (c *Cell) unassignVal() {
	c.val = 0
	for lc := range c.LinkedCells.iterate() {
		logUnconstraint(c, lc, c.val)
		lc.removeConstraint(c)
	}
}

func (c *Cell) constrain(val int, lc *Cell) {
	c.Constraints.addConstraint(val, lc)
}

func (c *Cell) removeConstraint(lc *Cell) {
	c.Constraints.removeConstraintByCell(lc)
}
