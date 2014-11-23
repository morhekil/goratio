package collector

import "github.com/morhekil/goratio/feeder/data"

// New does new
func New(c chan *data.Event) Reader {
	r := Reader{c: c}
	r.connect()
	r.prepare()
	return r
}
