package main

import "testing"

func TestCoordinate(t *testing.T) {
	c, err := newCoordinate(1, 2)
	if err != nil {
		t.Error(err)
	}
	if c.row != uint32(1) || c.col != uint32(2) {
		t.Error("Struct Coordinate looks wrong, expected 1,2 got:", c)
	}
}

func TestCoordinateRange(t *testing.T) {
	_, err := newCoordinate(20, 30)
	if err.Error() != "row out of range: 20" {
		t.Error("err was not nil got", err)
	}
	_, err = newCoordinate(2, 30)
	if err.Error() != "col out of range: 30" {
		t.Error("err was not nil got", err)
	}
	// if c.row != uint32(1) || c.col != uint32(2) {
	// 	t.Error("Struct Coordinate looks wrong, expected 1,2 got:", c)
	// }
}
