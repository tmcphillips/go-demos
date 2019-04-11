package main

import (
	"encoding/json"
	"fmt"
)

func newChannelFromSlice(values []int) <-chan int {
	channel := make(chan int, len(values))
	for _, value := range values {
		channel <- value
	}
	close(channel)
	return channel
}

func newSliceFromChannel(channel <-chan int) []int {
	size := len(channel)
	var values = make([]int, size)
	for i := 0; i < size; i++ {
		values[i] = <-channel
	}
	return values
}

func compareJSONOfSlices(expected []int, actual []int) string {
	actualJSON, _ := json.Marshal(actual)
	expectedJSON, _ := json.Marshal(expected)

	actualString := string(actualJSON)
	expectedString := string(expectedJSON)

	if actualString == expectedString {
		return ""
	}

	return fmt.Sprintf("\nExpect: %s\nActual: %s", expectedString, actualString)
}
