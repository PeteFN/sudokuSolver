package board

import "sudokuSolver/common"

type CellList struct {
	nextElementPos int
	length         int
	list           []*Cell
}

func newCellList() *CellList {
	return &CellList{
		nextElementPos: 0,
		length:         1,
		list:           make([]*Cell, 1),
	}
}

func (cl *CellList) add(c *Cell) {
	if cl.nextElementPos == cl.length {
		cl.list = append(cl.list, c)
		cl.length++
	} else {
		cl.list[cl.nextElementPos] = c
	}
	cl.nextElementPos++
}

func (cl *CellList) remove() (*Cell, error) {
	if cl.nextElementPos == 0 {
		return nil, common.NewListOutOfBoundsError(-1, 0)
	}
	cl.nextElementPos--
	c := cl.list[cl.nextElementPos]
	return c, nil
}

func (cl *CellList) size() int {
	return cl.nextElementPos
}

func (cl *CellList) get(pos int) (*Cell, error) {
	if pos < cl.length {
		return cl.list[pos], nil
	}
	return nil, common.NewListOutOfBoundsError(pos, cl.length)
}

func (cl *CellList) iterate() []*Cell {
	return cl.list
}
