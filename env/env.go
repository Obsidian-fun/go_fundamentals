/****
Source : https://dev.to/craicoverflow/a-no-nonsense-guide-to-environment-variables-in-go-a2f

****/

package main

import (
	"os"
	"fmt"

	"github.com/joho/godotenv"
)

// init is invoked before main,
func init() {
	// loads .env from the system,
	if err := godotenv.Load("../../../.env"); err != nil {
		fmt.Println("no .env file found");
	}
}

func main() {
	fmt.Println(os.Getenv("PASSWORD"));

	os.Setenv("Username", "Gigi Amore");
	fmt.Println(os.Getenv("Username"));
}


