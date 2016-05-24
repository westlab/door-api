package manager

import (
	"github.com/westlab/door-api/context"
	"github.com/westlab/door-api/job"
	"github.com/westlab/door-api/model"
)

const (
	htmlAnalyzerIDKey = "htmlAnalyzerID"
)

// HTMLAnalyzeManager is manager for job's HTMLAnalyzer
type HTMLAnalyzeManager struct {
	ctx          *context.Context
	htmlAnalyzer *job.HTMLAnalyzer
	idx          int64
}

// Start starts HTMLAnalyzeManager
func (manager *HTMLAnalyzeManager) Start() {
	go manager.run()
}

// run is job for html analyzer
func (manager *HTMLAnalyzeManager) run() {
	var b model.Browsing
	var browsings []model.Browsing

	for {
		browsings = model.GetBrowsingAfterID(manager.idx, 1000, false)
		for _, b := range browsings {
			manager.htmlAnalyzer.Manage(&b)
		}
		model.CreateOrUpdateMeta(htmlAnalyzerIDKey, string(b.ID))
		manager.idx = b.ID
	}
}

// NewHTMLAnalyzerManager creates HTMLAnalyzeManager
func NewHTMLAnalyzerManager(ctx *context.Context) *HTMLAnalyzeManager {
	htmlAnalyzer := job.NewHTMLAnalyzer()

	var idx int64
	var err error

	meta := model.SelectSingleMeta(htmlAnalyzerIDKey)
	if meta == nil {
		idx = 1
	} else {
		idx, err = meta.ToInt()
		if err != nil {
			idx = 1
		}
	}

	return &HTMLAnalyzeManager{ctx, htmlAnalyzer, idx}
}
