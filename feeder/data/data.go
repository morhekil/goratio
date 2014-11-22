package data

import "time"

type Event struct {
	User      string
	Domain    string
	Action    string
	Timestamp time.Time
}
