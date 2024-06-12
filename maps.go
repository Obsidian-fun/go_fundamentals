
package main

import (
	"fmt"
)

func main() {

	var mapLit map[string]int;
	var mapAssigned map[string]int;

	mapLit = map[string]int{"one":1,"two":2,"three":3};
	mapCreated := make(map[string]float32);
	mapAssigned = mapLit;

	mapCreated["Key1"] = 4.5;
	mapCreated["Key2"] = 3.141592653;

	fmt.Printf("Map literal at \"one\" is : %d\n", mapLit["one"]);
	fmt.Printf("Map Assigned at \"two\" is : %d\n", mapAssigned["two"]);
}


