package collector

import (
	"database/sql"
	"log"
	"os"

	"github.com/morhekil/goratio/feeder/data"
)

// fetch next row from the query results, and create page struct of it
func fetch(xs *sql.Rows) *data.Entry {
	d, err := scan(xs)
	if err != nil {
		log.Fatal(err)
	}
	return d
}

type reader struct {
	db     *sql.DB
	c      chan *data.Entry
	stmt   *sql.Stmt
	lastID int
	Count  int64
}

// Push loads the next batch of data and pushes it down the channel
func (r *reader) Push() {
	xs, err := r.stmt.Query(r.lastID)
	if err != nil {
		log.Fatal(err)
	}
	defer xs.Close()

	for xs.Next() {
		p := fetch(xs)
		r.publish(p)
	}

	if err := xs.Err(); err != nil {
		log.Fatal(err)
	}
}

// Close destroys the database connection
func (r *reader) Close() {
	r.db.Close()
}

func (r *reader) connect() {
	c, err := sql.Open("mysql", os.Getenv("DB"))
	if err != nil {
		log.Fatal(err)
	}
	r.db = c
}

func (r *reader) prepare() {
	s, err := r.db.Prepare(query())
	if err != nil {
		log.Fatal(err)
	}
	r.stmt = s
}

func (r *reader) publish(d *data.Entry) {
	r.lastID = d.ID
	r.Count++
	r.c <- d
}
