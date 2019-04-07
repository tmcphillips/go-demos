package main

import (
	"os"
	"strings"
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
	os.Args = strings.Fields("hello")
	main()
	// Output:
	// Hello World
}

func Example_main_NameOption() {
	os.Args = strings.Fields("hello -name Tim")
	main()
	// Output:
	// Hello Tim
}

func Example_main_GreetingOption() {
	os.Args = strings.Fields("hello -greeting Goodbye")
	main()
	// Output:
	// Goodbye World
}

func Example_main_NameAndGreetingOptions() {
	os.Args = strings.Fields("hello -name Tim -greeting Goodbye")
	main()
	// Output:
	// Goodbye Tim
}
