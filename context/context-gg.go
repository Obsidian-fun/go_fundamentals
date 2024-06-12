
package main

import (
	"context"
	"time"
	"log"
	"fmt"
)

func getUserData(ctx context.Context, userID int) (int, error) {

	ctx, cancel := context.WithCancel(ctx, time.Millisecond*200);
	defer cancel();

	val, err := tooSlow();
	if err != nil {
		return 0, err;
	}

	return val, nil;
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
