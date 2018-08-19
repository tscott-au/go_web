package main

import (
	"fmt"
)

const rangeMin = 1
const rangeMax = 10

// Coordiante building block of board
type Coordiante struct {
	row uint32
	col uint32
}

func newCoordinate(row, col uint32) (*Coordiante, error) {
	if !valid(row) {
		return nil, fmt.Errorf("row out of range: %v", row)
	}
	if !valid(col) {
		return nil, fmt.Errorf("col out of range: %v", col)
	}
	return &Coordiante{row: row, col: col}, nil
}

func valid(i uint32) bool {
	if i >= rangeMin && i <= rangeMax {
		return true
	}
	return false
}
