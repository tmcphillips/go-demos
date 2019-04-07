package hamming

import (
	"fmt"
)

func SequenceSource(valueArray []int, valueChannel chan int) {
	for _, value := range valueArray {
		valueChannel <- value
	}
	close(valueChannel)
}

func ExampleMultiplier() {

	chans := make([]chan int, 2)

	chans[0] = make(chan int)
	chans[1] = make(chan int)

	ch1 := chans[0] //make(chan int)
	ch2 := chans[1] // make(chan int)
	go SequenceSource([]int{1, 2, 3, 4, 5}, ch1)
	go Multiplier(ch1, ch2, 2)

	for n := range ch2 {
		fmt.Println(n)
	}

	// Output:
	// 2
	// 4
	// 6
	// 8
	// 10
}
