package util

type Average struct {
	values []int
}

func (a *Average) Append(v int) {
	a.values = append(a.values, v)
}

func (a *Average) Value() float64 {
	var sum float64

	for i := range a.values {
		sum += float64(a.values[i])
	}

	return sum / float64(len(a.values))
}
