package hamming

import (
	"fmt"
	"sync"
)

// LowPassFilter is an actor that outputs its integer inputs if they are below or equal to a configurable maximum.
func LowPassIntegerFilter(inputs chan int, outputs chan int, maximum int, waitgroup *sync.WaitGroup) {
	defer waitgroup.Done()
	for n := range inputs {
		if n <= maximum {
			outputs <- n
		}
	}
	close(outputs)
}

// IntegerMultiplier is an actor that outputs its integer inputs multiplied by a configurable constant factor.
func IntegerMultiplier(inputs chan int, outputs chan int, factor int, waitgroup *sync.WaitGroup) {
	defer waitgroup.Done()
	for n := range inputs {
		product := n * factor
		outputs <- product
	}
	close(outputs)
}

// IntegerPrinter is an actor that writes input integers to standard output.
func IntegerPrinter(inputs chan int, waitgroup *sync.WaitGroup) {
	defer waitgroup.Done()
	for n := range inputs {
		fmt.Println(n)
	}
}
