/***
Implentation of insertion sort.
Fastest for small arrays.
For larger arrays use heap or quick sort.

***/
package main

import (
	"fmt"
)

func sort(arr []int) []int {
	var j int;
	var key int;
	var buffer int; // holds a temporary value
	
	for i:=1 ; i<len(arr); i++ {
		key = arr[i];
		j = i - 1;
		for ; j >= 0; j--{
			if key < arr[j] {
				buffer = arr[j+1];
				arr[j+1] = arr[j];
				arr[j] = buffer;
			} 
		}		
	}
	return arr;
}

func main() {
	s := []int{43,22,1,12,67};
	a := []int{4,3,1,2,69};

	fmt.Println((sort(s)));
	fmt.Println((sort(a)));
}


