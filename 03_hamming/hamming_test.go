package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"sync"
	"testing"
)

func ExampleIntegerDistributor() {

	var waitgroup sync.WaitGroup
	waitgroup.Add(1)

	in := make(chan int, 5)
	out := []chan int{
		make(chan int, 5),
		make(chan int, 5),
	}

	go IntegerDistributor(in, out, &waitgroup)

	for _, value := range []int{1, 2, 3} {
		in <- value
	}
	close(in)

	for _, o := range out {
		for n := range o {
			fmt.Println(n)
		}
	}

	waitgroup.Wait()

	// Output:
	// 1
	// 2
	// 3
	// 1
	// 2
	// 3
}

func ExampleLowPassIntegerFilter() {

	var waitgroup sync.WaitGroup
	waitgroup.Add(1)

	in := make(chan int, 5)
	out := make(chan int, 5)

	go LowPassIntegerFilter(in, out, 5, &waitgroup)

	for _, value := range []int{1, 2, 3, 4, 5, 6, 7, 8} {
		in <- value
	}
	close(in)

	for n := range out {
		fmt.Println(n)
	}

	waitgroup.Wait()

	// Output:
	// 1
	// 2
	// 3
	// 4
	// 5
}

func ExampleIntegerMultiplier() {

	var waitgroup sync.WaitGroup
	waitgroup.Add(1)

	in := make(chan int, 5)
	out := make(chan int, 5)
	go IntegerMultiplier(in, out, 2, &waitgroup)

	for _, value := range []int{1, 2, 3, 4, 5} {
		in <- value
	}
	close(in)

	for n := range out {
		fmt.Println(n)
	}

	waitgroup.Wait()

	// Output:
	// 2
	// 4
	// 6
	// 8
	// 10
}

func ExampleIntegerPrinter() {

	var waitgroup sync.WaitGroup
	waitgroup.Add(1)

	in := make(chan int, 5)
	go IntegerPrinter(in, "\n", &waitgroup)

	for _, value := range []int{1, 2, 3, 4, 5} {
		in <- value
	}
	close(in)
	waitgroup.Wait()

	// Output:
	// 1
	// 2
	// 3
	// 4
	// 5
}

// func ExampleIntegerStreamMerge() {

// 	var waitgroup sync.WaitGroup
// 	waitgroup.Add(1)

// 	aInputs := make(chan int, 5)
// 	bInputs := make(chan int, 5)
// 	cOutputs := make(chan int, 100)

// 	go IntegerStreamMerge(aInputs, bInputs, cOutputs, &waitgroup)

// 	for _, value := range []int{1, 3, 5, 7, 9} {
// 		aInputs <- value
// 		bInputs <- value + 1
// 	}

// 	close(aInputs)
// 	close(bInputs)

// 	waitgroup.Wait()

// 	for n := range cOutputs {
// 		fmt.Println(n)
// 	}

// Output:
// 1
// 2
// 3
// 4
// 5
// 6
// 7
// 8
// 9
// 10
// }

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

func checkExpectedSlice(expected []int, actual []int) string {
	actualJSON, _ := json.Marshal(actual)
	expectedJSON, _ := json.Marshal(expected)

	actualString := string(actualJSON)
	expectedString := string(expectedJSON)

	if actualString == expectedString {
		return ""
	}

	return fmt.Sprintf("\nExpect: %s\nActual: %s", expectedString, actualString)
}

func TestIntegerStreamMerge(t *testing.T) {

	var table = []struct {
		description string
		aValues     []int
		bValues     []int
		expected    []int
	}{
		{"Both input channels empty",
			[]int{},
			[]int{},
			[]int{}},
		{"First input channel empty",
			[]int{},
			[]int{2, 4, 6, 8, 10},
			[]int{2, 4, 6, 8, 10}},
		{"Second input channel empty",
			[]int{1, 3, 5, 7, 9},
			[]int{},
			[]int{1, 3, 5, 7, 9}},
		{"Equal input sizes with alternating values",
			[]int{1, 3, 5, 7, 9},
			[]int{2, 4, 6, 8, 10},
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
		{"First input channel has single duplicate",
			[]int{1, 3, 5, 5, 7, 9},
			[]int{2, 4, 6, 8, 10},
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
		{"Second input channel has multiple duplicates",
			[]int{1, 3, 5, 7, 9},
			[]int{2, 2, 4, 4, 4, 6, 8, 8, 8, 8, 10},
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
		{"Two channels have duplicates between them",
			[]int{1, 3, 4, 5, 7, 9},
			[]int{2, 4, 6, 8, 9, 10},
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
	}

	for i, entry := range table {

		inChannelA := newChannelFromSlice(entry.aValues)
		inChannelB := newChannelFromSlice(entry.bValues)
		outChannelC := make(chan int, len(inChannelA)+len(inChannelB))

		var waitgroup sync.WaitGroup
		waitgroup.Add(1)
		go IntegerStreamMerge(inChannelA, inChannelB, outChannelC, &waitgroup)
		waitgroup.Wait()

		checkResult := checkExpectedSlice(
			entry.expected,
			newSliceFromChannel(outChannelC))

		if checkResult != "" {
			t.Error("[ Entry", i, "-", entry.description, "]", checkResult)
		}
	}
}

func Example_main_DefaultMax() {
	os.Args = []string{"hamming"}
	main()
	// Output:
	// 1, 2, 3, 4, 5, 6, 8, 9, 10, 12, 15, 16, 18, 20
}

func Example_main_Max5() {
	os.Args = strings.Fields("hamming -max 5")
	main()
	// Output:
	// 1, 2, 3, 4, 5
}

func Example_main_Max25() {
	os.Args = strings.Fields("hamming -max 25")
	main()
	// Output:
	// 1, 2, 3, 4, 5, 6, 8, 9, 10, 12, 15, 16, 18, 20, 24, 25
}

func Example_main_BadMaxOption() {
	os.Args = strings.Fields("hamming -max foo")
	main()
	// Output:
	// Could not convert max value to integer: strconv.Atoi: parsing "foo": invalid syntax
}

func Example_main_Max5_CommaSeparated() {
	os.Args = strings.Fields("hamming -max 5 -sep ,")
	main()
	// Output:
	// 1,2,3,4,5
}

func Example_main_Max5_NewlineSeparated() {
	os.Args = strings.Fields("hamming -max 5 -sep \\n")
	main()
	// Output:
	// 1
	// 2
	// 3
	// 4
	// 5
}

func Example_main_Max5_CommaTwoSpaceSeparated() {
	os.Args = []string{"hamming", "-max", "5", "-sep", ",  "}
	main()
	// Output:
	// 1,  2,  3,  4,  5
}

func Example_main_MaxOptionWithNoValue() {
	os.Args = []string{"hamming", "-max"}
	main()
	// Output:
}
