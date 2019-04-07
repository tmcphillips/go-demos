package hamming

import (
	"fmt"
)

func ExamplePrinter() {

	in := make(chan int, 1)
	go Printer(in)

	for _, value := range []int{1, 2, 3, 4, 5} {
		in <- value
	}
	close(in)

	// Output:
	// 1
	// 2
	// 3
	// 4
	// 5
}

func ExampleMultiplier() {

	in := make(chan int, 5)
	out := make(chan int, 5)
	go Multiplier(in, out, 2)

	for _, value := range []int{1, 2, 3, 4, 5} {
		in <- value
	}
	close(in)

	for n := range out {
		fmt.Println(n)
	}

	// Output:
	// 2
	// 4
	// 6
	// 8
	// 10
}
