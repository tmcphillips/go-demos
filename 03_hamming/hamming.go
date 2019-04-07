package hamming

func Multiplier(inputs chan int, outputs chan int, factor int) {
	for n := range inputs {
		product := n * factor
		outputs <- product
	}
	close(outputs)
}
