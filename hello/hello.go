package main

import (
	"fmt"
	"os"
)

// Program that outputs a simple greeting
func main() {

	name := "World"

	if len(os.Args) > 1 {
		name = os.Args[1]
	}

	fmt.Println(getDefaultGreeting(), name)
}

// getDefaultName returns the default name
func getDefaultName() string {
	return "World"
}

// getDefaultGreeting returns the default greeting
func getDefaultGreeting() string {
	return "Hello"
}
