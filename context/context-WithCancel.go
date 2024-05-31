/***
The context.WithCancel decorator function gets a context and returns another
context and a function called cancel . The returned context will be a copy of
the context that has a different done channel (the channel that marks that the
current context is done) that gets closed when the parent context does or
when the cancel function is called – whatever happens first.

***/
package main

import (
	"fmt"
	"time"
	"context"
)

func main() {

	ctx, cancel := context.WithCancel(context.Background());
	time.AfterFunc(time.Second*5, cancel);
	done := ctx.Done();

// Notice how case<-done gets executed after 5 seconds,
	for i:=0; ; i++ {
		select {
			case <-done:
				return
			case <-time.After(time.Second):
				fmt.Println("Tick ", i);
		}
	}
}



