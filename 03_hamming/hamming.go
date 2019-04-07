package hamming

import (
	"fmt"
	"sync"
)

func Multiplier(inputs chan int, outputs chan int, factor int, waitgroup *sync.WaitGroup) {
	defer waitgroup.Done()
	for n := range inputs {
		product := n * factor
		outputs <- product
	}
	close(outputs)
}

func Printer(inputs chan int, waitgroup *sync.WaitGroup) {
	defer waitgroup.Done()
	for n := range inputs {
		fmt.Println(n)
	}
}
