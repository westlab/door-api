package job

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBasic(t *testing.T) {
	soc := "/tmp/gotest"
	ch := make(chan string, 100)

	d := NewDoorReciever(soc, ch)
	// hot to stop server?
	// go d.Start()

}
