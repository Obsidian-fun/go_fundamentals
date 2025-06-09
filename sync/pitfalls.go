/***
Common pitfalls to using sync.Pool
#1 Type assetion overhead:
using raw interface{} can be error prone compared to using generics.
#2 Storing large objects
only use for frequently use small to medium sized objects.
for large objects use github.com/fatih/pool

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

