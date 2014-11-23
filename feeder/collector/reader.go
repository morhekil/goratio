package collector

import (
	"database/sql"
	"log"
	"os"

	"github.com/morhekil/goratio/feeder/data"
)

// fetch next row from the query results, and create page struct of it
func fetch(xs *sql.Rows) (uint64, *data.Event) {
	id, x, err := scan(xs)
	if err != nil {
		log.Fatal(err)
	}
	return id, x
}

type Reader struct {
	db     *sql.DB
	c      chan *data.Event
	stmt   *sql.Stmt
	lastID uint64
	Count  uint64
}

// Push loads the next batch of data and pushes it down the channel
func (r *Reader) Push() {
	xs, err := r.stmt.Query(r.lastID)
	if err != nil {
		log.Fatal(err)
	}
	defer xs.Close()

	for xs.Next() {
		id, p := fetch(xs)
		r.publish(id, p)
	}

	if err := xs.Err(); err != nil {
		log.Fatal(err)
	}
}

// Close destroys the database connection
func (r *Reader) Close() {
	r.db.Close()
}

func (r *Reader) connect() {
	c, err := sql.Open("mysql", os.Getenv("DB"))
	if err != nil {
		log.Fatal(err)
	}
	r.db = c
}

func (r *Reader) prepare() {
	s, err := r.db.Prepare(query())
	if err != nil {
		log.Fatal(err)
	}
	r.stmt = s
}

func (r *Reader) publish(id uint64, d *data.Event) {
	r.lastID = id
	r.Count++
	r.c <- d
}
