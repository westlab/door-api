package common

import (
	"math"
	"time"
)

// TimeTuple is a tule which holds time.Time
type TimeTuple struct {
	X, Y time.Time
}

// ZipTime performs python like zip
func ZipTime(x, y []time.Time) []TimeTuple {
	num := int64(math.Min(float64(len(x)), float64(len(y))))
	r := make([]TimeTuple, num, num)
	for i := 0; i < int(num); i++ {
		r[i] = TimeTuple{x[i], y[i]}
	}
	return r
}

// PairwiseTime makes a pair from given sequence
func PairwiseTime(t []time.Time) []TimeTuple {
	x := make([]time.Time, len(t))
	copy(x, t)
	// delete first element
	x = x[1:]
	return ZipTime(t, x)
}
