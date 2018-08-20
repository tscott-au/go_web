package main

import (
	"reflect"

	"github.com/satori/go.uuid"
)

// Island hold coordinates of each island and keeps a track of hits
type Island struct {
	location CoordinateMap
	hits     CoordinateMap
	id       uuid.UUID
}

// Forrested - checks if all coordinates in an island have been guessed
func (i *Island) Forrested() bool {
	return reflect.DeepEqual(i.hits, i.location)
}

// NewIsland make a new island
func NewIsland(x *Coordinate, fn func() *Coordinates) (r *Island, err error) {

	offsets := (*fn())
	for i := 0; i < len(offsets); i++ {
		offsets[i][Row] = offsets[i][Row] + x[Row]
		offsets[i][Col] = offsets[i][Col] + x[Col]

	}

	l, err := NewCoordinateMap(&offsets)
	if err != nil {
		return
	}

	var empty Coordinates
	h, err := NewCoordinateMap(&empty)
	if err != nil {
		return
	}

	return &Island{location: *l, hits: *h, id: uuid.Must(uuid.NewV4())}, nil
}

// Atoll - make an atoll shape
func Atoll() *Coordinates { return &Coordinates{{0, 0}, {0, 1}, {1, 1}, {2, 0}, {2, 1}} }

// Dot - dot co-ordinates
func Dot() *Coordinates { return &Coordinates{{0, 0}} }

// Lshape - make an lshape
func Lshape() *Coordinates { return &Coordinates{{0, 0}, {1, 0}, {2, 0}, {2, 1}} }

// Sshape - make an sshape
func Sshape() *Coordinates { return &Coordinates{{0, 1}, {0, 2}, {1, 0}, {1, 1}} }

// Square - make a square
func Square() *Coordinates { return &Coordinates{{0, 0}, {0, 1}, {1, 0}, {1, 1}} }
