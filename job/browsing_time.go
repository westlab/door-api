package job

import (
	"time"

	"github.com/westlab/door-api/model"
)

// BrowsingTimer calculrate browsing time.
// Browsing time is the difference between request.
type BrowsingTimer struct {
	repository map[string]*model.Browsing
	gcTime     time.Time
	gcDuration time.Duration
	timeout    time.Duration
}

// NewBrowsingTimer create a BrowsingTimeCalculator instance
//
// size: hint for memory allocation of map
// timeout: threshold for deleting Browsing instance from repository
// gcDuration: interval for gc
func NewBrowsingTimer(size int64, timeout int64, gcDuration int64) *BrowsingTimer {
	if gcDuration < 0 {
		gcDuration = 0
	}
	if timeout < 0 {
		timeout = 30 * 60
	}
	b := make(map[string]*model.Browsing, size)
	return &BrowsingTimer{b, time.Now(),
		time.Duration(timeout), time.Duration(gcDuration)}
}

// Add adds Browsing to BrowsingTimer
func (browsingTimer *BrowsingTimer) Add(b *model.Browsing) {
	k := b.SrcIP
	existedBrowsing, ok := browsingTimer.repository[k]
	if !ok {
		browsingTimer.repository[k] = b
		return
	}

	previousTime := existedBrowsing.Timestamp.Time
	newTime := b.Timestamp.Time
	browsingTime := int64(previousTime.Sub(newTime) / time.Second)
	existedBrowsing.BrowsingTime = browsingTime
	existedBrowsing.Update()
	browsingTimer.repository[k] = b

	if time.Since(browsingTimer.gcTime) > browsingTimer.gcDuration*time.Second {
		browsingTimer.gcRepository()
		browsingTimer.gcTime = time.Now()
	}
}

// gcRepository perform gc
func (browsingTimer *BrowsingTimer) gcRepository() {
	now := time.Now()
	gcPoint := now.Add(-browsingTimer.timeout * time.Second)
	for key, b := range browsingTimer.repository {
		if b.Timestamp.Time.Before(gcPoint) {
			delete(browsingTimer.repository, key)
		}
	}
}
