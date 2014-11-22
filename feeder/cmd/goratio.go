package main

import (
	"time"

	"github.com/morhekil/goratio/feeder/collector"
	"github.com/morhekil/goratio/feeder/data"
	"github.com/morhekil/goratio/feeder/emitter"
)

func main() {
	d := make(chan *data.Entry)
	defer close(d)

	r := collector.New(d)
	defer c.Close()

	go emitter.Pull(d)
	for {
		c.Push()
		time.Sleep(1 * time.Second)
	}
}
