package main

import (
	"fmt"
	"sync"
)

// IntegerDistributor is an actor that outputs its integer inputs on each provided output channel.
func IntegerDistributor(inputs <-chan int, outputChannels []chan int, waitgroup *sync.WaitGroup) {

	defer waitgroup.Done()

	for n := range inputs {
		for _, outputChannel := range outputChannels {
			outputChannel <- n
		}
	}

	for _, outputChannel := range outputChannels {
		close(outputChannel)
	}
}

// LowPassIntegerFilter is an actor that outputs its integer inputs if they are below or equal to a configurable maximum.
func LowPassIntegerFilter(inputs <-chan int, outputs chan<- int, maximum int, waitgroup *sync.WaitGroup) {

	defer waitgroup.Done()

	for n := range inputs {
		if n <= maximum {
			outputs <- n
		} else {
			break
		}
	}

	close(outputs)

	for range inputs {
	}
}

// IntegerMultiplier is an actor that outputs its integer inputs multiplied by a configurable constant factor.
func IntegerMultiplier(inputs <-chan int, outputs chan<- int, factor int, waitgroup *sync.WaitGroup) {

	defer waitgroup.Done()

	for n := range inputs {
		product := n * factor
		outputs <- product
	}

	close(outputs)
}

// IntegerPrinter is an actor that writes input integers to standard output.
func IntegerPrinter(inputs <-chan int, waitgroup *sync.WaitGroup) {

	defer waitgroup.Done()

	for n := range inputs {
		fmt.Println(n)
	}
}

// IntegerStreamMerge is actor that merges two, ordered integer streams
func IntegerStreamMerge(inputA <-chan int, inputB <-chan int, outputC chan<- int, waitgroup *sync.WaitGroup) {

	defer waitgroup.Done()

	var a, b, c, lastC int
	haveA, haveB := false, false
	aOpen, bOpen := true, true

	for aOpen || bOpen {

		if !haveA {
			a, aOpen = <-inputA
			if aOpen {
				// fmt.Println("a:", a)
				haveA = true
			}
		}

		if !haveB {
			b, bOpen = <-inputB
			if bOpen {
				// fmt.Println("b:", b)
				haveB = true
			}
		}

		if !haveA && !haveB {
			break
		}

		if !haveA {
			c = b
			haveB = false
		} else if !haveB {
			c = a
			haveA = false
		} else if a < b {
			c = a
			haveA = false
		} else if b < a {
			c = b
			haveB = false
		} else if a == b {
			c = a
			haveA = false
			haveB = false
		}

		if c != lastC {
			outputC <- c
			lastC = c
		}
	}

	close(outputC)
}
