/***
using sync.Pool in HFT as a light example

***/

package main

import (
	"sync"
)

type Order struct {
	ID string;
	Amount float64;
}

var orderPool = sync.Pool {
	New: func() interface{} {
		return &Order{}
	},
}

// in the request handler,
func HandleProcess() {
	order := orderPool.Get().(*Order);
	defer orderPool.Put(order);

}


func main() {

}

