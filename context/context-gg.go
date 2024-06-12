
package main

import (
	"fmt"
	"context"
	"net/http"
	"time"
)

func getUserData(ctx contex.Context, userID int) (int, error) {
	val, err := tooSlow();
	if err != nil {
		return 0, err;
	}

	return val, nil;
}

func tooSlow() (int , error)){
	time.Sleep(time.Millisecond*500);

	return 666, nil;
}

func main() {

	ctx := context.Background();
	val, err := getUserData(ctx, )

}
