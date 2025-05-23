package bitarray

import "fmt"

type OutOfBoundsError struct {
	length byte
	index  byte
}

func (o OutOfBoundsError) Error() string { return fmt.Sprintf("Bit array index is out of bounds (array length: %d, index: %d)", o.length, o.index) }

func (ba *BitArray) CheckBoundsError(index byte) bool {
	return index >= ba.length
}