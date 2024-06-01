
package main

import (
	"fmt"
	"context"
	"time"
)

func main() {
	// Setting the timeout as 5 seconds from now,
	ctx, cancel := context.WithTimeout(context.Background(),time.Second*5);
	time.AfterFunc(time.Second*10,cancel);
	done := ctx.Done();

	for i:=0; ; i++ {
		select {
			case <-done:
				fmt.Println("Err: ",ctx.Err());
				return
			case <-time.After(time.Second*1):
				fmt.Println("tick : ",i);
		}
	}
}

