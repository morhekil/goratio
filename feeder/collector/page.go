package collector

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql" // mysql driver for database/sql

	"github.com/morhekil/goratio/feeder/data"
)

// query to use when fetching the next batch of data
func query() string {
	return "SELECT id, user_id, url, controller, action, server_addr, " +
		"http_method, form_data FROM page_logs WHERE id > ? LIMIT 1000"
}

func scan(xs *sql.Rows) (*data.Entry, error) {
	p := page{}
	err := xs.Scan(&p.ID, &p.UserID, &p.URL, &p.Controller, &p.Action,
		&p.Server, &p.Method, &p.FormData)
	if err != nil {
		return nil, err
	}
	return p.dataEntry(), nil
}

// page structure describes a single website activity event
type page struct {
	ID         int
	UserID     sql.NullInt64
	URL        string
	Controller string
	Action     string
	Server     string
	Method     string
	FormData   sql.NullString
}

func (p *page) dataEntry() *data.Entry {
	return &data.Entry{
		ID:         p.ID,
		UserID:     int(p.UserID.Int64),
		Server:     p.Server,
		Controller: p.Controller,
		Action:     p.Action,
		Method:     p.Method,
	}
}
