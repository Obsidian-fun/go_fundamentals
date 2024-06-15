/***
Exercise 14.5 from "The way to Go" by Ivo Balbaert.
Program : add 2 numbers and wait for return

***/

package main


// Size of channel buffer,
const N = 2;

type Empty interface{};
type semaphore chan Empty;

func (s semaphore) addAndWait(a int, b int) int{

	var e = new(Empty);
	var sum = 0;
	for i:=0; i<N; i++ {
		sum = a+b;
		s<-e;
	}
	return sum;

}

func main() {

	sem := make(semaphore, N);
	go sem.addAndWait(4,6);


}

