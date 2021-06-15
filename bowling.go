package main

type Scorer struct {
	Throws []int
}

func (s *Scorer) AddThrow(numberOfPinsDown int) {

	s.Throws = append(s.Throws, numberOfPinsDown)
}
