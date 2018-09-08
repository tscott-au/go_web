package game

import (
	"fmt"
)

// Row index for Coordinate array
const Row = 0

// Col index for Coordinate array
const Col = 1

// Coordinate building block of board
type Coordinate [2]uint8

// Coordinates is an array of coordinates
type Coordinates []Coordinate

// CoordinateMap is a unique set of co-ordinates.
type CoordinateMap map[Coordinate]bool

// NewCoordinate make a new coordinate
func NewCoordinate(c ...uint8) (*Coordinate, error) {
	if len(c) != 2 {
		return nil, fmt.Errorf("wrong number of input, expected 2 ints got: %v", c)
	}

	return &Coordinate{c[Row], c[Col]}, nil
}

// NewCoordinateMap make a new coordinates
func NewCoordinateMap(c *Coordinates) (*CoordinateMap, error) {
	m := make(CoordinateMap)

	for i, v := range *c {
		//v := v
		if m[v] {
			return nil, fmt.Errorf("duplicate detected key[%d] = value %v", i, v)
		}
		m[v] = true
	}

	return &m, nil
}
