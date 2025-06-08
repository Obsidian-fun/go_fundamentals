/***
from, https://reliasoftware.com/blog/golang-sync-pool

sync.Pool uses local pools per processor to minimize lock contention.
each goroutine accesses
***/

package main

import (
	"fmt"
	"sync"
	"encoding/json"
)

var encoderPool = sync.Pool {
		New: func() interface{} {
					return json.NewEncoder;
		},
	}

func GetEncoder(w io.Writer) *json.Encoder {
	enc := encoderPool.Get().(*json.Encoder);
	enc.Reset(w);
	return enc
}

func ReturnEncoder(enc *json.Encoder) {
	enc.Reset(nil); // reset state
	encoderPool.Put(enc);
}

func main() {

}



