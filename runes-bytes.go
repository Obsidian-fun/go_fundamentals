
package main

import (
	"fmt"
)

func main(){

	const a = `⌘`;

	fmt.Printf("Plain string %s\n", a);
	fmt.Printf("Quoted string %q\n", a);

	fmt.Printf("Hex bytes: ");
	for i:=0; i<len(a); i++ {
		fmt.Printf("%x",a[i]);
	}
	fmt.Println();
}


