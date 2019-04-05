package main

import (
	"flag"
	"fmt"
	"os"
)

// Program that outputs a simple greeting
func main() {

	var commandLine = flag.NewFlagSet("", 0)
	var name = commandLine.String("name", getDefaultName(), "Name of someone to greet")
	var greeting = commandLine.String("greeting", getDefaultGreeting(), "Greeting to use")
	commandLine.Parse(os.Args[1:])

	fmt.Println(*greeting, *name)
}

// getDefaultName returns the default name
func getDefaultName() string {
	return "World"
}

// getDefaultGreeting returns the default greeting
func getDefaultGreeting() string {
	return "Hello"
}
