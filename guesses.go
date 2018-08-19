package main

type guessMap map[Coordiante]*Coordiante

// Guesses - hold hit / miss for guesses
type Guesses struct {
	hits   guessMap
	misses guessMap
}

func newGuesses() *Guesses {
	return &Guesses{hits: make(guessMap), misses: make(guessMap)}
}
