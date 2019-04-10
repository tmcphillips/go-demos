package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"sync"
)

const defaultMaxOption = "20"

// Generates all of the Hamming numbers up to a maximum value.
func main() {

	var commandLine = flag.NewFlagSet("", 0)
	var maxValueOption = commandLine.String("max", defaultMaxOption, "Maximum Hamming number to generate")
	commandLine.Parse(os.Args[1:])

	maxValue, err := strconv.Atoi(*maxValueOption)
	if err != nil {
		fmt.Println("Could not convert max value to integer:", err)
		return
	}

	const channelSize = 100

	valuesToMultiplyBy2 := make(chan int, channelSize)
	valuesToMultiplyBy3 := make(chan int, channelSize)
	valuesToMultiplyBy5 := make(chan int, channelSize)
	multipliedBy2Values := make(chan int, channelSize)
	multipliedBy3Values := make(chan int, channelSize)
	multipliedBy5Values := make(chan int, channelSize)
	merged2x3xValues := make(chan int, channelSize)
	merged2x3x5xValues := make(chan int, channelSize)
	filteredValues := make(chan int, channelSize)
	valuesToPrint := make(chan int, channelSize)

	var waitgroup sync.WaitGroup
	waitgroup.Add(8)

	go IntegerMultiplier(valuesToMultiplyBy2, multipliedBy2Values, 2, &waitgroup)
	go IntegerMultiplier(valuesToMultiplyBy3, multipliedBy3Values, 3, &waitgroup)
	go IntegerMultiplier(valuesToMultiplyBy5, multipliedBy5Values, 5, &waitgroup)
	go IntegerStreamMerge(multipliedBy2Values, multipliedBy3Values, merged2x3xValues, &waitgroup)
	go IntegerStreamMerge(merged2x3xValues, multipliedBy5Values, merged2x3x5xValues, &waitgroup)
	go LowPassIntegerFilter(merged2x3x5xValues, filteredValues, maxValue, &waitgroup)
	go IntegerDistributor(filteredValues, []chan int{valuesToMultiplyBy2, valuesToMultiplyBy3, valuesToMultiplyBy5, valuesToPrint}, &waitgroup)
	go IntegerPrinter(valuesToPrint, &waitgroup)

	filteredValues <- 1

	waitgroup.Wait()
}

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
