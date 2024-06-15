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

const N = 5;
sem := make(semaphore, N);

fmt.Println(reflect.TypeOf(sem));

}
