package job

import (
	"github.com/westlab/door-api/context"
	"github.com/westlab/door-api/job"
	"github.com/westlab/door-api/model"
)

const (
	browsingIDKey = "browsingID"
)

// BrowsingTimeManager manages browsing time
type BrowsingTimeManager struct {
	cxt           *context.Context
	browsingTimer *job.BrowsingTimer
	idx           int64
}

// Start starts BrowsingTimeManager
func (manager *BrowsingTimeManager) Start() {
	go manager.run()
}

// run is job for browsing timer
func (manager *BrowsingTimeManager) run() {
	var b model.Browsing
	var browsings []model.Browsing

	for {
		browsings = model.GetBrowsingAfterID(manager.idx, 1000, true)
		for _, b = range browsings {
			manager.browsingTimer.Add(&b)
		}
		model.CreateOrUpdateMeta(browsingIDKey, string(b.ID))
		manager.idx = b.ID
	}
}

// NewBrowsingTimer creates BrowsingTimeManager
func NewBrowsingTimeManager(cxt *context.Context) *BrowsingTimeManager {
	browsingTimer := job.NewBrowsingTimer(10000, 60*10, 60*10)
	var idx int64

	meta := model.SelectSingleMeta(browsingIDKey)
	if meta != nil {
		idx = 0
	} else {
		idx, err := m.ToInt()
		if err != nil {
			idx = 0
		}
	}

	return &BrowsingTimeManager{cxt, browsingTimer, idx}
}
