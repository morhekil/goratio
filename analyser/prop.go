package analyser

import (
	"fmt"

	"github.com/morhekil/goratio/analyser/timeframe"
)

// PropEvent is a concrete implementation of Prop, based on Events
type PropEvent struct {
	name string
	r    *Repository
}

// Analyse performs the analysis of the prop over the given timeframe
func (p PropEvent) Analyse(t timeframe.Moment) {
	fmt.Printf("%s\t%+v\n", p.name, t)
}
