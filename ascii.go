// ASCII value is as simple as int(<<letter>>)

package main

import "fmt"

func main() {

	letter := 'A';
	slice  := []string{"ball","cat"};

	fmt.Printf("\nASCII value of 'A' : %d\n\n", int(letter));
	fmt.Printf("ASCII value of 'a' : %d\n\n", int('a'));
	fmt.Printf("ASCII value of 'b' : %d\n\n", int(slice[0][0]));
	fmt.Printf("ASCII value of 'c' : %d\n\n", int(slice[1][0]));
}
