package main

import (
	"os"
	"testing"
)

func TestGetDefaultGreeting(t *testing.T) {
	result := getDefaultGreeting()
	if result != "Hello" {
		t.Error("Got the wrong default greeting")
	}
}

func Example_main_NoArguments() {
	os.Args = []string{""}
	main()
	// Output:
	// Hello World
}

func Example_main_OneArgument() {
	os.Args = []string{"", "Tim"}
	main()
	// Output:
	// Hello Tim
}
