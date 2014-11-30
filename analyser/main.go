package main

import (
	"github.com/morhekil/goratio/analyser/core"
	"github.com/morhekil/goratio/analyser/elastic"
	"github.com/morhekil/goratio/analyser/timeframe"
)

func main() {
	r := elastic.Repository{}
	c := core.New(r)
	t := timeframe.Now()
	c.Process(t)

	// m := timeframe.Now()
	// s := stats.Calc(m)
	// es := s.Events(m)
	// for e := range es {
	// }
}
