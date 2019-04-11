package main

import (
	"fmt"
	"sync"
	"testing"
)

func TestIntegerDistributor(t *testing.T) {

	inputValues := []int{1, 2, 3}
	inChannel := newChannelFromSlice(inputValues)
	outChannels := []chan int{
		make(chan int, 5),
		make(chan int, 5),
	}

	var waitgroup sync.WaitGroup
	waitgroup.Add(1)
	go IntegerDistributor(inChannel, outChannels, &waitgroup)
	waitgroup.Wait()

	expectedOutputs := inputValues
	for i := range []int{0, 1} {
		diff := compareJSONOfSlices(
			expectedOutputs,
			newSliceFromChannel(outChannels[i]))
		if diff != "" {
			t.Error("[ Out channel", i, "]", diff)
		}
	}
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

		diff := compareJSONOfSlices(
			entry.expected,
			newSliceFromChannel(outChannelC))

		if diff != "" {
			t.Error("[ Entry", i, "-", entry.description, "]", diff)
		}
	}
}
