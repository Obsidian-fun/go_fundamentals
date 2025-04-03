/***
Struct padding: Go compilers may pad some bytes
after certain fields of struct values. The padded bytes are counted for struct sizes. So the size of a
struct type may be not a simple sum of the sizes of all its fields.

Go Optimizations 101, pg 12

***/

package main

import (
	"fmt"

	"unsafe"
)

type T1 struct {
	a int32; // 4 bytes
	b string; // 16 bytes
	c *string; // 8 bytes 
}


func main() {

	fmt.Println(unsafe.Sizeof(T1{})); // returns 32 bytes instead of 28 , padded 4 extra bytes

}


