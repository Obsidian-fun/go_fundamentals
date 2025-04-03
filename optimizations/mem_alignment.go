/***
Memory alignment :

***/

package main

import (

	"unsafe"
)

type T struct {
	a int32;
}


func main() {
	print(unsafe.Alignof(T{}));
	print("\n");
	print(unsafe.Sizeof(T{}));
	print("\n");
}

