package job

import (
	"log"
	"net"
)

type DoorReciver struct {
	unixSocket          string
	toHTTPReconstructor chan<- *string
}

func NewDoorReciever(unixSocket string, toHTTPReconstructor chan<- *string) *DoorReciver {
	return &DoorReciver{unixSocket, toHTTPReconstructor}
}

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
		toHTTPReconstructor <- &string(data)
	}
}

func (d *DoorReciver) Start() {
	l, err := net.Listen("unix", d.unixSocket)
	if err != nil {
		log.Println(err)
	}
	defer l.Close()

	for {
		fd, err := l.Accept()
		if err != nil {
			log.Printf(err)
		}

		go d.receive(fd)
	}
}
