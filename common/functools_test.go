package common

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestZipTime(t *testing.T) {
	now := time.Now()
	s1 := make([]time.Time, 10)
	s2 := make([]time.Time, len(s1))
	for i := 0; i < len(s1); i++ {
		s1[i] = now.Add(time.Duration(i) * time.Second)
	}
	copy(s2, s1)
	timeZip := ZipTime(s1, s2)

	assert.Equal(t, 10, len(timeZip))
	assert.Equal(t, timeZip[0].x, timeZip[0].y)
	assert.Equal(t, timeZip[9].x, timeZip[9].y)

	s3 := make([]time.Time, 5)
	for i := 0; i < len(s3); i++ {
		s3[i] = s2[i]
	}
	timeZip = ZipTime(s2, s3)

	assert.Equal(t, 5, len(timeZip))
	assert.Equal(t, timeZip[0].x, timeZip[0].y)
	assert.Equal(t, timeZip[4].x, timeZip[4].y)
}

func TestPairwiseTime(t *testing.T) {
	now := time.Now()
	s1 := make([]time.Time, 10)
	for i := 0; i < len(s1); i++ {
		s1[i] = now.Add(time.Duration(i) * time.Second)
	}

	timePair := PairwiseTime(s1)
	assert.Equal(t, 9, len(timePair))
	assert.Equal(t, timePair[0].y, timePair[1].x)
}
