package main

import (
	"reflect"
	"testing"
)

func TestIslandAtoll(t *testing.T) {
	x, err := NewCoordinate(1, 1)
	if err != nil {
		t.Error(err)
	}
	i, err := NewIsland(x, Atoll)
	if err != nil {
		t.Error("failed NewIsland() ", err)
	}

	if i == nil {
		t.Error("didn't make an island:", i)
	}
	//r := fmt.Sprintf("%v", i)
	//s := "&{map[[1 1]:true [1 2]:true [2 2]:true [3 1]:true [3 2]:true] map[]}"
	s, err := NewCoordinateMap(&Coordinates{{1, 1}, {1, 2}, {2, 2}, {3, 1}, {3, 2}})

	if err != nil {
		t.Error(err)
	}
	if !reflect.DeepEqual((*s), i.location) {
		t.Errorf("Atoll island wrong shape.  Expected 1: got 2: \n1: %v\n2: %v", (*s), i.location)
	}

}

func TestIslandSshape(t *testing.T) {
	x, err := NewCoordinate(2, 2)
	if err != nil {
		t.Error(err)
	}
	i, err := NewIsland(x, Sshape)
	if err != nil {
		t.Error("failed NewIsland() ", err)
	}

	if i == nil {
		t.Error("didn't make an island:", i)
	}

	s, err := NewCoordinateMap(&Coordinates{{2, 4}, {3, 2}, {3, 3}, {2, 3}})
	if err != nil {
		t.Error(err)
	}

	//r := fmt.Sprintf("%v", i)
	//s := "&{map[[2 4]:true [3 2]:true [3 3]:true [2 3]:true] map[]}"
	if !reflect.DeepEqual((*s), i.location) {
		t.Errorf("Sshape island wrong shape. Expected 1: got 2: \n1: %v\n2: %v", (*s), i.location)
	}

}

func TestIslandDotForrest(t *testing.T) {
	x, err := NewCoordinate(1, 1)
	if err != nil {
		t.Error(err)
	}
	i, err := NewIsland(x, Dot)
	if err != nil {
		t.Error("failed NewIsland() ", err)
	}

	if i == nil {
		t.Error("didn't make an island:", i)
	}

	s, err := NewCoordinateMap(&Coordinates{{1, 1}})
	if err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual((*s), i.location) {
		t.Errorf("Dot island wrong shape. Expected 1: got 2: \n1: %v\n2: %v", (*s), i.location)
	}

	if i.Forrested() {
		t.Errorf("Dot island forrested error location != hits but got %v", i)
	}

	for k := range *s {
		i.hits[k] = true
	}

	if !i.Forrested() {
		t.Errorf("Dot island should be forrested but got %v", i)
	}
}

func TestIslandSquareForrest(t *testing.T) {
	x, err := NewCoordinate(9, 9)
	if err != nil {
		t.Error(err)
	}
	i, err := NewIsland(x, Square)
	if err != nil {
		t.Error("failed NewIsland() ", err)
	}

	if i == nil {
		t.Error("didn't make an island:", i)
	}

	s, err := NewCoordinateMap(&Coordinates{{9, 9}, {9, 10}, {10, 9}, {10, 10}})
	if err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual((*s), i.location) {
		t.Errorf("Square island wrong shape. Expected 1: got 2: \n1: %v\n2: %v", (*s), i.location)
	}

	if i.Forrested() {
		t.Errorf("Square island forrested error location != hits but got %v", i)
	}

	skip, y := 0, 1
	for k := range *s {
		if skip >= y {
			i.hits[k] = true
		}
		skip++
	}

	if i.Forrested() {
		t.Errorf("Square island should NOT be forrested but got %v", i)
	}

	for k := range *s {
		i.hits[k] = true
	}

	if !i.Forrested() {
		t.Errorf("Square island should NOT be forrested but got %v", i)
	}
}

func IslandHelper(t *testing.T, s IslandShape, c ...uint8) (i *Island, ok bool) {
	x, err := NewCoordinate(c...)
	if err != nil {
		t.Error(err)
		ok = false
	}

	i, err = NewIsland(x, s)
	if err != nil {
		t.Error("failed NewIsland() ", err)
		ok = false
	}

	if i == nil {
		t.Error("didn't make an island:", i)
		ok = false
	}
	ok = true

	return

}

func TestIslandSquareOverlaps(t *testing.T) {
	i, ok := IslandHelper(t, Square, 1, 3)
	if !ok {
		return
	}

	j, ok := IslandHelper(t, Square, 2, 3)
	if !ok {
		return
	}

	if !i.Overlaps(j) && !j.Overlaps(i) {
		t.Errorf("Island overlap problem")
	}

	k, ok := IslandHelper(t, Square, 3, 3)
	if !ok {
		return
	}

	if i.Overlaps(k) {
		t.Errorf("Island overlap problem2")
	}

	m, ok := IslandHelper(t, Lshape, 1, 3)
	if !ok {
		return
	}

	if !k.Overlaps(m) {
		t.Errorf("Island overlap problem3")
	}
}
