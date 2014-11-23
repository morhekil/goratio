package emitter

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	elastigo "github.com/mattbaird/elastigo/lib"

	"github.com/morhekil/goratio/feeder/data"
)

const index = "goratio"
const doctype = "goratio-event"

var es *elastigo.Conn

// Setup emitter
func Setup() {
	es = elastigo.NewConn()
	host := os.Getenv("ESHOST")
	if host != "" {
		es.SetHosts([]string{host})
	}
}

// Pull accepts a channel generated by a collector, and emits
// it into the destination storage
func Pull(c chan *data.Event) {
	for e := range c {
		e.Hydrate()
		// debug(e)
		elasticise(e)
	}
}

func indexName(e *data.Event) string {
	return fmt.Sprintf("goratio-%4d.%02d.01",
		e.Timestamp.Year(), e.Timestamp.Month())
}

func elasticise(e *data.Event) {
	args := map[string]interface{}{"timestamp": e.Timestamp.String()}
	_, err := es.Index(indexName(e), doctype, e.ID, args, e)
	if err != nil {
		log.Fatal(err)
	}
}

func debug(e *data.Event) {
	s, err := json.Marshal(e)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", s)
}
