package synth

import "math"

const Pi = 3.1415926535897932384626433

type Sin struct {
	index      int
	frequency  float64
	sampleRate float64
}

func (s *Sin) Read() float64 {
	angle := s.frequency * float64(s.index) / s.sampleRate
	value := math.Sin(2.0 * Pi * angle)

	s.index += 1

	return value
}

func NewSin(frequency, sampleRate float64) *Sin {
	return &Sin{
		index:      0,
		frequency:  frequency,
		sampleRate: sampleRate,
	}
}
