package game

import (
	"errors"
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

// Overlaps check to see if new island will overlap an existing one
func (i *Island) Overlaps(newi *Island) bool {
	for k := range i.location {
		if newi.location[k] {
			return true
		}
	}
	return false
}

// NewIsland make a new island
func NewIsland(x *Coordinate, s IslandShape) (r *Island, err error) {

	var offsets Coordinates
	switch s {
	case Atoll:
		offsets = Coordinates{{0, 0}, {0, 1}, {1, 1}, {2, 0}, {2, 1}}
	case Dot:
		offsets = Coordinates{{0, 0}}
	case Lshape:
		offsets = Coordinates{{0, 0}, {1, 0}, {2, 0}, {2, 1}}
	case Sshape:
		offsets = Coordinates{{0, 1}, {0, 2}, {1, 0}, {1, 1}}
	case Square:
		offsets = Coordinates{{0, 0}, {0, 1}, {1, 0}, {1, 1}}
	default:
		return nil, errors.New("unsupported IslandShape type")

	}

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

// IslandShape - supported shape types
type IslandShape int

const (
	// Atoll - make an atoll shape
	Atoll IslandShape = iota
	// Dot - just a single point
	Dot
	// Lshape - make an lshape
	Lshape
	// Sshape - make an sshape
	Sshape
	// Square - make a square
	Square
)
