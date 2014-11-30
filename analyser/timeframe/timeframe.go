package timeframe

import "time"

// Moment interface represents a timeframe for the analyser
type Moment interface {
}

// Hourly represents statistical timeframe, with hourly precision
type Hourly struct {
	time.Time
}

// Now returns a timeframe corresponding to the current time
func Now() Moment {
	return Hourly{}
}
