package analyser

import "github.com/morhekil/goratio/analyser/timeframe"

// Prop is a property of the system, that is happening over time as an event
type Prop interface {
	Analyse(timeframe.Moment)
}

// Repository provides units with access to the event data
type Repository interface {
	Props() []Prop
}

// New analysing unit is returned, connected to the given repository
// for data access
func New(r Repository) unit {
	return unit{repo: &r}
}

type unit struct {
	repo *Repository
}

// Process event data for the given moment of the log
func (u unit) Process(t timeframe.Moment) {
	for _, p := range (*u.repo).Props() {
		p.Analyse(t)
	}
}
