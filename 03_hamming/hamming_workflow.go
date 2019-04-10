package main

import (
	"sync"
)

func main() {

	const maxValue int = 10
	const channelSize = 100

	var waitgroup sync.WaitGroup
	waitgroup.Add(8)

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