package main

import (
	"log"
	"time"

	"github.com/morhekil/goratio/feeder/collector"
	"github.com/morhekil/goratio/feeder/data"
	"github.com/morhekil/goratio/feeder/emitter"
)

func stats(r *collector.Reader) {
	t := time.NewTicker(time.Second * 5)

	go func() {
		n := uint64(0)
		for _ = range t.C {
			if r.Count > n {
				log.Printf("Processed %d records", r.Count-n)
				n = r.Count
			}
		}
	}()
}
func main() {
	d := make(chan *data.Event)
	defer close(d)

	r := collector.New(d)
	defer r.Close()

	go emitter.Pull(d)
	stats(&r)

	for {
		r.Push()
		time.Sleep(1 * time.Second)
	}
}
