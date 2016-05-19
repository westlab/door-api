package job

import (
	"strconv"
	"strings"
	"time"

	"github.com/westlab/door-api/model"
)

// HTTPReconstructor is for reconstructing HTTP from packet
type HTTPReconstructor struct {
	repository    map[string]*model.HTTPCommunication
	gcTime        time.Time
	gcDuration    time.Duration
	fromReciverCh <-chan *string
}

// Start starts HTTPReconstructor
func (h *HTTPReconstructor) Start() {
	for {
		data := <-h.fromReciverCh
		dpi := convertStrToDPI(data)
		h.add(dpi)
	}
}

// add adds DPI to resitory to maintain HTTP connection
func (h *HTTPReconstructor) add(dpi *model.DPI) {
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

// convertStrToDPI converts string to DPI
// SrcIP, DstIP, SrcPort, DstPort, Timestamp, Rule, data
func convertStrToDPI(data *string) (dpi *model.DPI) {
	d := strings.SplitN(*data, ",", 7)
	srcPort, _ := strconv.Atoi(d[2])
	dstPort, _ := strconv.Atoi(d[3])
	timestamp, _ := time.Parse(d[4], "2006-01-02 15:04:06")

	dpi = &model.DPI{
		SrcIP:     d[0],
		DstIP:     d[1],
		SrcPort:   int64(srcPort),
		DstPort:   int64(dstPort),
		Timestamp: timestamp,
		Rule:      d[5],
		Data:      d[6],
	}
	return dpi
}

// NewHTTPReconstructor creates a new HTTPReconstructor
func NewHTTPReconstructor(size int64, gcDuration int64, fromReciverCh <-chan *string) *HTTPReconstructor {
	if gcDuration < 0 {
		gcDuration = 0
	}
	m := make(map[string]*model.HTTPCommunication, size)
	return &HTTPReconstructor{m, time.Now(), time.Duration(gcDuration), fromReciverCh}
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
