package common

import "fmt"

type ListOutOfBoundsError struct {
	requestedIndex int
	actualLength   int
}

func NewListOutOfBoundsError(requested int, actual int) error {
	return &ListOutOfBoundsError{
		requestedIndex: requested,
		actualLength:   actual,
	}
}

func (l ListOutOfBoundsError) Error() string {
	return fmt.Sprintf("Array out of bounds at index=%d, length=%d", l.requestedIndex, l.actualLength)
}
