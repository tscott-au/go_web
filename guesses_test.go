package main

import "testing"

func TestGuesses(t *testing.T) {
	c1, err := newCoordinate(1, 2)
	if err != nil {
		t.Error(err)
	}
	m := newGuesses()

	if m == nil {
		t.Error("wanted Guesses got", m)
	}

	m.misses[*c1] = c1
	if len(m.misses) != 1 {
		t.Error("misses not 1")
	}

	m.misses[*c1] = c1
	if len(m.misses) > 1 {
		t.Error("misses should be 1")
	}

	c2, err := newCoordinate(1, 2)
	if err != nil {
		t.Error(err)
	}

	m.misses[*c2] = c2
	if len(m.misses) > 1 {
		t.Error("misses should be 1")
	}

	c3, err := newCoordinate(2, 2)
	if err != nil {
		t.Error(err)
	}

	m.misses[*c3] = c3
	if len(m.misses) != 2 {
		t.Error("misses should be 2")
	}

}
