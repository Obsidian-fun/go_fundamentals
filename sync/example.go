/***
from, https://reliasoftware.com/blog/golang-sync-pool


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
		}

	}

func main() {
	


}



