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

	ctx, cancel := context.withCancel(context.Background());
	time.After(time.Second*5, done);
	done := ctx.Done();

