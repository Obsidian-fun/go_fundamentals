/***
sources, 
https://stackoverflow.com/questions/10493411/what-is-bit-masking
https://lorbic.com/bitmasking/

the act of applying a mask to a value. Accomplished by doing,
AND : extract a subset of bits in the value
OR  : set a subset of bits in the value
XOR : toggle a set of bits in the value

Imagine we have a bunch of characterisitics such as 
Properties  Name Bit
Property A     0      0
Property B     1      1<<0
Property C     2      1<<1
Property D     3      1<<2

***/

package main

import (
	"fmt"
)

func main() {
	// first set desired properties using OR,
	mask := 0;
	mask |= 1 << 1;  // Enable property B
	mask |= 1 << 3;  // Enable property D
	fmt.Println(mask);

	// then extract the properties using AND,
	for i:=0; i<32; i++ {
		if mask & (1<<i) != 0 {
			fmt.Println("bit", i , "is set");
		}
	}

	// we can also toggle bits, using XOR,so it works great with booleans
	mask = 10; // so it has bit on in 1st and 3rd position
	mask ^= 1<<3; // toggle third position off, value is now 2
	fmt.Println(mask);


}
/***
USAGE :
Good for:

    Feature flags
    User roles/permissions
    Categorical flags with a known upper limit (< 64 ideally)
		Options maybe ? (long/ short, call/put, buy/sell)

Not great for:

    Huge dynamic lists
    Highly descriptive data (bitmasks are not self-documenting)
    Frequent schema changes (bit positions are hard-coded)
***/





