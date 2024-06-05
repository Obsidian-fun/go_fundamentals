
package main

import (
	"fmt"
	"bytes"
)

func main() {

	var buf = bytes.NewBuffer(make([]byte, 26));

	var texts = []string{
				`This is a sentence,`,
				`What you see here, is also a sentence..`,
				`And yet another! Another sentence..`,
				`Guess what? You guessed it, a sentence.`,
				}

	for i:= range texts {
		buf.Reset();
		buf.WriteString(texts[i]);
		fmt.Println("Length",buf.Len(), "\t capacity", buf.Cap());

	}
}
