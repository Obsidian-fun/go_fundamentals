/***
from, https://reliasoftware.com/blog/golang-sync-pool

sync.Pool uses local pools per processor to minimize lock contention.
each goroutine accesses
***/

package main

import (
	"io"
	"sync"
	"encoding/json"
)

var encoderPool = sync.Pool {
		New: func() interface{} {
					return json.NewEncoder(nil);
		},
	}

func GetEncoder(w io.Writer) *json.Encoder {
	enc := encoderPool.Get().(*json.Encoder);
	return enc
}

func ReturnEncoder(enc *json.Encoder) {
	encoderPool.Put(enc);
}

func main() {

}



