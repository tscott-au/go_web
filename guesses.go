package main

// Guesses - hold hit / miss for guesses
type Guesses struct {
	hits   CoordinateMap
	misses CoordinateMap
}

// NewGuesses make a new guesses struct
func NewGuesses() *Guesses {
	return &Guesses{hits: make(CoordinateMap), misses: make(CoordinateMap)}
}

func (g *Guesses) hit(c *Coordinate) {
	g.hits[*c] = true
}

func (g *Guesses) miss(c *Coordinate) {
	g.misses[*c] = true
}
