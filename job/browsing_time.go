package job

import (
	"time"

	"github.com/westlab/door-api/model"
)

// BrowsingTimeCalculator calculrate browsing time.
// Browsing time is the difference between request.
type BrowsingTimeCalculator struct {
	repository map[string]*model.Browsing
	gcTime     time.Time
	gcDuration time.Duration
	timeout    time.Duration
}

// NewBrowsingTimeCalculator create a BrowsingTimeCalculator instance
//
// size: hint for memory allocation of map
// timeout: threshold for deleting Browsing instance from repository
// gcDuration: interval for gc
func NewBrowsingTimeCalculator(size int64, timeout int64, gcDuration int64) *BrowsingTimeCalculator {
	if gcDuration < 0 {
		gcDuration = 0
	}
	if timeout < 0 {
		timeout = 30 * 60
	}
	b := make(map[string]*model.Browsing, size)
	return &BrowsingTimeCalculator{b, time.Now(),
		time.Duration(timeout), time.Duration(gcDuration)}
}

// Add adds Browsing to BrowsingTimeCalculator
func (browsingTimeCal *BrowsingTimeCalculator) Add(b *model.Browsing) {
	k := b.SrcIP
	existedBrowsing, ok := browsingTimeCal.repository[k]
	if !ok {
		browsingTimeCal.repository[k] = b
		return
	}

	previousTime := existedBrowsing.Timestamp.Time
	newTime := b.Timestamp.Time
	browsingTime := int64(previousTime.Sub(newTime) / time.Second)
	existedBrowsing.BrowsingTime = browsingTime
	existedBrowsing.Update()
	browsingTimeCal.repository[k] = b

	if time.Since(browsingTimeCal.gcTime) > browsingTimeCal.gcDuration*time.Second {
		browsingTimeCal.gcRepository()
		browsingTimeCal.gcTime = time.Now()
	}
}

// gcRepository perform gc
func (browsingTimeCal *BrowsingTimeCalculator) gcRepository() {
	now := time.Now()
	gcPoint := now.Add(-browsingTimeCal.timeout * time.Second)
	for key, b := range browsingTimeCal.repository {
		if b.Timestamp.Time.Before(gcPoint) {
			delete(browsingTimeCal.repository, key)
		}
	}
}
