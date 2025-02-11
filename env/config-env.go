/***
wrapping environmental variables in a configuration package for long term ease of maintenance,
Offers :
	- Multiline struct storage for many APIs, passwords, etc.
	- Default values along with new values set in function

source : https://dev.to/craicoverflow/a-no-nonsense-guide-to-environment-variables-in-go-a2f
***/

package main

import (
	"os"
_	"fmt"

_	"github.com/joho/godotenv"
)

type Details struct {
	Username string;
	APIKey	string;
}

type Config struct {
	details Details;
}

// a new config struct per call,
func New() *Config {
	return &Config{
		details : Details {
			Username: os.Getenv("Username"),
			APIKey: os.Getenv("APIKey"),
		},
	}
}

// Helper function to return default values or 
func GetEnv(key string, default_value string) string{ 
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return default_value
}


func main() {}



