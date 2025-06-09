
package main

import (
	"sync"
)

type User struct {
	ID uint64;
	Data interface{};
}

var structPool = sync.Pool {
	New: func() interface{} {
		return &User{}
	},
}

func Process() {
	user :=	structPool.Get().(*User);
	structPool.Put(user);
// reset fields if required
}

func main() {

}

