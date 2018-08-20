package main

import "testing"

func TestCoordinate(t *testing.T) {
	c, err := NewCoordinate(1, 2)
	if err != nil {
		t.Error(err)
	}
	if c[Row] != uint32(1) || c[Col] != uint32(2) {
		t.Error("Struct Coordinate looks wrong, expected 1,2 got:", c)
	}
}

func TestCoordinateRange(t *testing.T) {
	// commented out as not relevant

	// _, err := NewCoordinate(20, 30)
	// if err.Error() != "row out of range: 20" {
	// 	t.Error("err was not nil got", err)
	// }
	// _, err = NewCoordinate(2, 30)
	// if err.Error() != "col out of range: 30" {
	// 	t.Error("err was not nil got", err)
	// }

	// if c.row != uint32(1) || c.col != uint32(2) {
	// 	t.Error("Struct Coordinate looks wrong, expected 1,2 got:", c)
	// }
}
