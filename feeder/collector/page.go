package collector

import (
	"database/sql"
	"fmt"
	"regexp"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql" // mysql driver for database/sql

	"github.com/morhekil/goratio/feeder/data"
)

var domainRx = regexp.MustCompile(`^www\d*\.`)
var dataRx = regexp.MustCompile(`^--- '(.+)'$`)
var localtime, _ = time.LoadLocation("Local")

const isotime = "2006-01-02 15:04:05"

// query to use when fetching the next batch of data
func query() string {
	return "SELECT id, user_id, url, controller, action, server_addr, " +
		"http_method, form_data, params, created_at " +
		"FROM page_logs WHERE id > ? LIMIT 1000"
}

func scan(xs *sql.Rows) (uint64, *data.Event, error) {
	var id uint64
	p := page{}
	err := xs.Scan(&id, &p.userID, &p.url, &p.controller, &p.action,
		&p.server, &p.method, &p.data, &p.params, &p.createdAt)
	if err != nil {
		return 0, nil, err
	}
	return id, p.event(), nil
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
	params     string
	createdAt  string
}

func (p *page) event() *data.Event {
	return &data.Event{
		User:      strconv.FormatInt(p.userID.Int64, 10),
		Domain:    p.domain(),
		Action:    p.description(),
		Timestamp: p.timestamp(),
	}
}

func (p *page) domain() string {
	return domainRx.ReplaceAllLiteralString(p.server, "")
}

func (p *page) description() string {
	s := fmt.Sprintf("%s %s/%s", p.method, p.controller, p.action)
	id := p.target(s)
	if id != "" {
		s += "/" + id
	}
	return s
}

func (p *page) timestamp() time.Time {
	t, _ := time.ParseInLocation(isotime, p.createdAt, localtime)
	return t
}

func (p *page) target(s string) string {
	f, ok := mappers[s]
	if ok {
		return f(p)
	}
	return ""
}
