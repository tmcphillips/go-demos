package hamming

import "fmt"

func Multiplier(inputs chan int, outputs chan int, factor int) {
	for n := range inputs {
		product := n * factor
		outputs <- product
	}
	close(outputs)
}

func Printer(inputs chan int) {
	for n := range inputs {
		fmt.Println(n)
	}
}
