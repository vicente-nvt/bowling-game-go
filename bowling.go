package main

type Scorer struct {
	Throws []int
	Frames []int
}

func (s *Scorer) AddThrow(numberOfPinsDown int) {
	s.Throws = append(s.Throws, numberOfPinsDown)

	throwsLen := len(s.Throws)
	if throwsLen > 0 && throwsLen%2 == 0 {
		frameThrows := s.Throws[throwsLen-1] + s.Throws[throwsLen-2]
		s.Frames = append(s.Frames, frameThrows)
	}
}

func (s *Scorer) GetFrameScore(index int) (int, error) {
	return 0, nil
}
