// ASCII value is as simple as int(<<letter>>)

package main

import "fmt"

func main() {

	letter := 'A';
	fmt.Printf("\nASCII value of 'A' : %d\n\n", int(letter));
	fmt.Printf("ASCII value of 'a' : %d\n\n", int('a'));
}
