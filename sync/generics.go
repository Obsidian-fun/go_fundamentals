/***
# Type assetion overhead:

using raw interface{} can be error prone compared to using generics.

***/

package main

import (
	"sync"
)

type Pool[T any] struct{
	Data sync.Pool;
}

func NewPool[T any](newFunc func() T) *Pool[T] {
	return &Pool[T] {
		Data: sync.Pool { 
			New: func() interface{} {
				return newFunc()
			},
		},
	}
}

func main() {}

