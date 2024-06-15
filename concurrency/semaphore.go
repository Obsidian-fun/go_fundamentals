/***

Implementing a semaphore using a buffered channel

***/

package main

import (
//	"fmt"
)


type Empty interface{};
type semaphore chan Empty;

const N = 5;

// produce empty resources
func (s semaphore) Produce(n int) {
	e := new(Empty);
	for i:=0; i<n; i++ {
		s<-e;
	}
}

// release those empty resources
func (s semaphore) Release(n int) {
	for i:=0; i<n; i++ {
		<-s;
	}
}

// mutex - rapid access of resources without blocking,
func (s semaphore) Lock(){
	s.Produce(1);  //locks in first resource
}

func (s semaphore) Unlock(){
	s.Release(1); // unlocks first resource
}

// Wait and signal...
func (s semaphore) Wait(n int) {
	s.Produce(n);
}

func (s semaphore) Signal() {
	s.Release(1);
}

func main() {
	sem := make(semaphore, N);
	go sem.Produce(N);
	go sem.Release(N);

}

