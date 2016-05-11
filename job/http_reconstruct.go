package job

import (
	"github.com/westlab/door-api/model"
	"time"
)

// HTTPReconstructor is for reconstructing HTTP from packet
type HTTPReconstructor struct {
	repository map[string]*model.HTTPCommunication
	gcTime     time.Time
	gcDuration time.Duration
}

// Add adds DPI to resitory to maintain HTTP connection
func (h *HTTPReconstructor) Add(dpi model.DPI) {
	k := dpi.ToKey()
	hc, ok := h.repository[k]
	if !ok {
		h.repository[k] = dpi.ToHTTPCommunication()
		return
	}

	data, _ := dpi.ParseData()
	switch dpi.Rule {
	case "GET":
		// GET should come first
		h.repository[k] = dpi.ToHTTPCommunication()
	case "Host:":
		hc.Host = data
	case "ContentType:":
		hc.ContentType = data
	case "<title":
		hc.Title = data
		if hc.IsValid("") {
			hc.ToBrowsing().Save()
			delete(h.repository, k)
		}
	}

	if time.Since(h.gcTime) > h.gcDuration*time.Second {
		h.gcRepository()
		h.gcTime = time.Now()
	}
}

// NewHTTPReconstructor creates a new HTTPReconstructor
func NewHTTPReconstructor(size int64, gcDuration int64) *HTTPReconstructor {
	if gcDuration < 0 {
		gcDuration = 0
	}
	m := make(map[string]*model.HTTPCommunication, size)
	return &HTTPReconstructor{m, time.Now(), time.Duration(gcDuration)}
}

// gcRepository perform govage collection
func (h *HTTPReconstructor) gcRepository() {
	// Lock
	now := time.Now()
	gcPoint := now.Add(-h.gcDuration * time.Second)
	for k, hc := range h.repository {
		if hc.Time.Before(gcPoint) {
			delete(h.repository, k)
		}
	}
}
