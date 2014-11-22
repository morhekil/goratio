package data

import (
	"crypto/md5"
	"fmt"
	"time"
)

// Event data respresentation
type Event struct {
	ID        string
	User      string
	Domain    string
	Action    string
	Timestamp time.Time
}

// Hydrate populates event record with additional derived data
func (e *Event) Hydrate() {
	s := fmt.Sprintf("%s-%s-%s", e.User, e.Action, e.Timestamp)
	e.ID = fmt.Sprintf("%x", md5.Sum([]byte(s)))
}
