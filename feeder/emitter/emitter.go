package emitter

import (
	"fmt"

	"github.com/morhekil/goratio/feeder/data"
)

// Pull accepts a channel generated by a collector, and emits
// it into the destination storage
func Pull(c chan *data.Event) {
	fmt.Printf("PULL")
	for entry := range c {
		fmt.Printf("%v\n", entry)
	}
}