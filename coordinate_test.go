package main

import "testing"

func TestCoordinate(t *testing.T) {
	c, err := NewCoordinate(1, 2)
	if err != nil {
		t.Error(err)
	}
	if c[Row] != uint8(1) || c[Col] != uint8(2) {
		t.Error("Struct Coordinate looks wrong, expected 1,2 got:", c)
	}
}

func TestCoordinateMap(t *testing.T) {

	c := &Coordinates{{1, 1}, {1, 2}}
	_, err := NewCoordinateMap(c)
	if err != nil {
		t.Error(err)
	}

	c = &Coordinates{{1, 1}, {1, 2}, {1, 1}}
	_, err = NewCoordinateMap(c)
	if err == nil {
		t.Error("Expected duplicate key error got:", err)
	}
}
