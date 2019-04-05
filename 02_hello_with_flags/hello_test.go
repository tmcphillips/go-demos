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

func TestGetDefaultName(t *testing.T) {
	result := getDefaultName()
	if result != "World" {
		t.Error("Got the wrong default name")
	}
}

func Example_main_NoArguments() {
	os.Args = []string{""}
	main()
	// Output:
	// Hello World
}

func Example_main_OneArgument() {
	os.Args = []string{"", "-name", "Tim"}
	main()
	// Output:
	// Hello Tim
}
