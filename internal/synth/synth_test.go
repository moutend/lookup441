package synth

import (
	"encoding/binary"
	"io/ioutil"
	"math"
	"testing"
)

func Test(t *testing.T) {
	hiA := NewSin(440.0, 44100.0)

	data := []byte{}

	for i := 0; i < 44100*2; i++ {
		value := float32(hiA.Read())
		b := make([]byte, 4)

		binary.LittleEndian.PutUint32(b, math.Float32bits(value))
		data = append(data, b...)
	}

	ioutil.WriteFile("output.raw", data, 0644)
}
