/***

Implementing a semaphore using a buffered channel

***/

package main

import (
	"fmt"
	"reflect"
)

func main() {

type Empty interface{};
type semaphore chan Empty;

fmt.Println(reflect.TypeOf(semaphore));

}
