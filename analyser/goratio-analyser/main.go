package main

import (
	"github.com/morhekil/goratio/analyser"
	"github.com/morhekil/goratio/analyser/elastic"
	"github.com/morhekil/goratio/analyser/timeframe"
)

func main() {
	r := elastic.Repository{}
	a := analyser.New(r)
	t := timeframe.Now()
	a.Process(t)

	// m := timeframe.Now()
	// s := stats.Calc(m)
	// es := s.Events(m)
	// for e := range es {
	// }
}
