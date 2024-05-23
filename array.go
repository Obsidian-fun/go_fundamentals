
// Addition using pointer to arrays [taken from 'effective go', increases efficiency but is distinct like in C,
// Also non-idiomatic as compared to slices.

package main

import (
	"fmt"
)

func Sum(a *[3]float64) (sum float64){
	for _, v := range *a {
		sum += v;
	}
	return sum;
}

func main() {
	array := [...]float64{3.14 ,2.78 ,0};
	fmt.Println(Sum(&array));
}



