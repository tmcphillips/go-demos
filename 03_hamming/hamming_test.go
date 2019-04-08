package hamming

import (
	"fmt"
	"sync"
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

	for _, value := range []int{1, 7, 2, 6, 19, 3, 6, 7, 4, 9, 5} {
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
	go IntegerPrinter(in, &waitgroup)

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

func ExampleIntegerStreamMerge() {

	var waitgroup sync.WaitGroup
	waitgroup.Add(1)

	aInputs := make(chan int, 5)
	bInputs := make(chan int, 5)
	cOutputs := make(chan int, 100)

	go IntegerStreamMerge(aInputs, bInputs, cOutputs, &waitgroup)

	for _, value := range []int{1, 3, 5, 7, 9} {
		aInputs <- value
		bInputs <- value + 1
	}

	close(aInputs)
	close(bInputs)

	waitgroup.Wait()

	for n := range cOutputs {
		fmt.Println(n)
	}

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
}
