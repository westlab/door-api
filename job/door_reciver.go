package job

import (
	"log"
	"net"
)

// DoorReciver recieve data from door through the unix domain socket
type DoorReciver struct {
	unixSocket          string
	toHTTPReconstructor chan<- *string
}

// NewDoorReciever creates DoorReciver
func NewDoorReciever(unixSocket string, toHTTPReconstructor chan<- *string) *DoorReciver {
	return &DoorReciver{unixSocket, toHTTPReconstructor}
}

// recieve is called when data arrives
func (d *DoorReciver) receive(c net.Conn) {
	defer c.Close()

	for {
		// ringbuffer may be better?
		buf := make([]byte, 1024)
		nr, err := c.Read(buf)
		if err != nil {
			log.Println(err)
			return
		}

		data := buf[0:nr]
		// TODO: do not convert data to string here.
		pdata := string(data)
		d.toHTTPReconstructor <- &pdata
	}
}

// Start starts DoorReciver
func (d *DoorReciver) Start() {
	l, err := net.Listen("unix", d.unixSocket)
	if err != nil {
		log.Println(err)
	}
	defer l.Close()

	for {
		fd, err := l.Accept()
		if err != nil {
			log.Println(err)
		}
		go d.receive(fd)
	}
}
