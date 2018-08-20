package main

import (
	"fmt"
)

const rangeMin = 1
const rangeMax = 10

// Row index for Coordinate array
const Row = 0

// Col index for Coordinate array
const Col = 1

// Coordinate building block of board
type Coordinate [2]uint32

// Coordinates is a group of coordinates
type Coordinates []Coordinate

// CoordinateMap used to enforce uniqueness in array of coordinates.
type CoordinateMap map[Coordinate]bool

// NewCoordinate make a new coordinate
func NewCoordinate(c ...uint32) (*Coordinate, error) {
	if len(c) != 2 {
		return nil, fmt.Errorf("wrong number of input, expected 2 ints got: %v", c)
	}
	// if !valid(c[Row]) {
	// 	return nil, fmt.Errorf("row out of range: %v", c[Row])
	// }
	// if !valid(c[Col]) {
	// 	return nil, fmt.Errorf("col out of range: %v", c[Col])
	// }
	return &Coordinate{c[Row], c[Col]}, nil
}

// NewCoordinateMap make a new coordinates
func NewCoordinateMap(c Coordinates) (*CoordinateMap, error) {
	m := make(CoordinateMap)

	for i, v := range c {
		//v := v
		if m[v] {
			return nil, fmt.Errorf("duplicate key at item %d of value %v", i, v)
		}
		m[v] = true
	}

	return &m, nil
}

// func valid(i uint32) bool {
// 	if i >= rangeMin && i <= rangeMax {
// 		return true
// 	}
// 	return false
// }
