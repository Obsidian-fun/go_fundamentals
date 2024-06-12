
package main

import (
	"context"
	"time"
	"log"
	"fmt"
)

type Response struct {
	value int;
	err		error;
}


func getUserData(ctx context.Context, userID int) (int, error) {

	ctx, cancel := context.WithTimeout(ctx, time.Millisecond*200);
	defer cancel();

	respch := make(chan Response);

	go func() {
		val, err := tooSlow();
		respch <- Response{
			value: val,
			err: err,
		}
	}()

	for {
		select {
			case <- ctx.Done():
				return 0, fmt.Errorf("fetching data from third party took too long");
			case resp := <-respch:
				return resp.value, resp.err;
		}
	}
}

func tooSlow() (int , error){
	time.Sleep(time.Millisecond*500);

	return 666, nil;
}

func main() {

	start := time.Now();
	ctx := context.Background();
	userID := 10;
	val, err := getUserData(ctx, userID);
	if err != nil {
		log.Fatal(err);
	}

	fmt.Println("result :", val);
	fmt.Println(time.Since(start));
}
