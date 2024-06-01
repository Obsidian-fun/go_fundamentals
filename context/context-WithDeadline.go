
package main

import (
	"context"
	"fmt"
	"time"
)

func main() {

// Deadline set to 5 seconds..
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(5*time.Second));

// cancel set to 10 seconds...5 secs more than deadline
	time.AfterFunc(time.Second*10, cancel);
	done := ctx.Done();

	for i:=0; ;i++ {
		select {
			case <-done:
				fmt.Println("exit: ",ctx.Err());
				return
			case <-time.After(time.Second):
				fmt.Println("tick ",i);
		}
	}
}



