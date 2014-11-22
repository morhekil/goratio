package collector

import (
	"database/sql"
	"fmt"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql" // mysql driver for database/sql

	"github.com/morhekil/goratio/feeder/data"
)

// query to use when fetching the next batch of data
func query() string {
	return "SELECT id, user_id, url, controller, action, server_addr, " +
		"http_method, form_data, created_at " +
		"FROM page_logs WHERE id > ? LIMIT 1000"
}

func scan(xs *sql.Rows) (uint64, *data.Event, error) {
	var id uint64
	p := page{}
	err := xs.Scan(&id, &p.userID, &p.url, &p.controller, &p.action,
		&p.server, &p.method, &p.data, &p.createdAt)
	if err != nil {
		return 0, nil, err
	}
	return id, p.dataEvent(), nil
}

// page structure describes a single website activity event
type page struct {
	userID     sql.NullInt64
	url        string
	controller string
	action     string
	server     string
	method     string
	data       sql.NullString
	createdAt  string
}

func (p *page) dataEvent() *data.Event {
	return &data.Event{
		User:      strconv.FormatInt(p.userID.Int64, 10),
		Domain:    p.normServer(),
		Action:    p.event(),
		Timestamp: p.timestamp(),
	}
}

func (p *page) normServer() string {
	return p.server
}

func (p *page) event() string {
	return fmt.Sprintf("%s %s/%s", p.method, p.controller, p.action)
}

func (p *page) timestamp() time.Time {
	const isotime = "2006-01-02 15:04:05"
	t, _ := time.Parse(isotime, p.createdAt)
	return t
}
