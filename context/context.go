package context

import (
	"github.com/westlab/door-api/conf"
)

var cxt *Context

// Context carrys shared information
type Context struct {
	conf       *conf.Config
	recieveChs []chan *string
}

// NewContext create Context
func NewContext(conf *conf.Config) *Context {
	if cxt != nil {
		return cxt
	}
	recieveChs := make([]chan *string, len(conf.Sockets), len(conf.Sockets))
	for i := 0; i < len(recieveChs); i++ {
		recieveChs[i] = make(chan *string, 1000000)
	}
	cxt = &Context{conf, recieveChs}
	return cxt
}

// GetContext returns Context instance
func GetContext() *Context {
	return cxt
}

// GetConf returns Conf
func (c *Context) GetConf() *conf.Config {
	return c.conf
}

// GetRecieverChs returns recieveChs
func (c *Context) GetRecieverChs() []chan *string {
	return c.recieveChs
}
