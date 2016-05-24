package job

import (
	"errors"
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
		dpi, err := convertStrToDPI(data)
		if err != nil {
			continue
		}
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
	case "Content-Type:":
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
func convertStrToDPI(data *string) (dpi *model.DPI, err error) {
	d := strings.SplitN(*data, ",", 7)
	if len(d) != 7 {
		return nil, errors.New("data format is wrong :" + *data)
	}
	srcPort, err := strconv.Atoi(d[2])
	if err != nil {
		return nil, err
	}
	dstPort, err := strconv.Atoi(d[3])
	if err != nil {
		return nil, err
	}
	loc, _ := time.LoadLocation("Asia/Tokyo")
	timestamp, err := time.ParseInLocation("2006-01-02 15:04:05", d[4], loc)
	if err != nil {
		return nil, err
	}
	rule, ok := mapIDtoRule(d[5])
	if ok != true {
		return nil, errors.New("Rule is not matched in rule list")
	}

	dpi = model.NewDPI(d[0], d[1], int64(srcPort), int64(dstPort), timestamp, rule, d[6])
	return dpi, nil
}

// mapIDtoRule maps rule id to rule
//
// 500:/HTTP/
// 501:/GET/
// 502:/POST/
// 503:/Host:/
// 504:/<title/
// 505:/Content-Length:/
// 506:/Content-Type:/
// 507:/Accept-Encoding:/
// 508:/Content-Encoding:/
// 509:/Status:/
func mapIDtoRule(id string) (rule string, ok bool) {
	// map is the easist way
	switch id {
	case "500":
		return "HTTP", true
	case "501":
		return "GET", true
	case "502":
		return "POST", true
	case "503":
		return "Host:", true
	case "504":
		return "<title", true
	case "505":
		return "Content-Length:", true
	case "506":
		return "Content-Type:", true
	case "507":
		return "Accept-Encoding:", true
	case "508":
		return "Content-Encoding:", true
	case "509":
		return "Status:", true
	default:
		return "", false
	}
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
