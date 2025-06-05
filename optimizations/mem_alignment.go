/***
Memory alignment :

***/

package main

import (

	"unsafe"
)

type T struct {
	a int32;
	b string;
	c complex64;
}

func main() {
	print(unsafe.Alignof(T{}));
	print("\n");
	print(unsafe.Sizeof(T{}));
	print("\n");
}

