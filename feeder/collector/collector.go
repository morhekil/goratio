package collector

import "github.com/morhekil/goratio/feeder/data"

// New does new
func New(c chan *data.Event) reader {
	r := reader{c: c}
	r.connect()
	r.prepare()
	return r
}
