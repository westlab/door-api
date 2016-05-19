package manager

import (
	"github.com/westlab/door-api/context"
	"github.com/westlab/door-api/job"
)

type HTTPJobManager struct {
	recievers      []*job.DoorReciver
	reconstructors []*job.HTTPReconstructor
}

func NewHTTPJobManager(cxt *context.Context) *HTTPJobManager {
	sockets := cxt.GetConf().Sockets
	recieverToReconstructorCh := cxt.GetRecieverChs()
	recievers := make([]*job.DoorReciver, len(sockets), len(sockets))
	reconstructors := make([]*job.HTTPReconstructor, len(sockets), len(sockets))

	for i := 0; i < len(sockets); i++ {
		recievers[i] = job.NewDoorReciever(sockets[i], recieverToReconstructorCh[i])
		reconstructors[i] = job.NewHTTPReconstructor(10000, 60*10, recieverToReconstructorCh[i])
	}
	return &HTTPJobManager{recievers, reconstructors}
}

func (h *HTTPJobManager) Start() {
	for _, r := range h.recievers {
		go r.Start()
	}

	for _, r := range h.reconstructors {
		go r.Start()
	}
}
