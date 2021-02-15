package util

type ZeroCross struct {
	last  float32
	count int
}

func (zc *ZeroCross) Apply(v float32) {
	if zc.last >= 0.0 && v < 0.0 || zc.last < 0.0 && v >= 0.0 {
		zc.count += 1
	}
	zc.last = v
}

func (zc *ZeroCross) Count() int {
	return zc.count
}
