package main

import (
	"os"
	"strings"
)

func Example_main_DefaultMax() {
	os.Args = []string{"hamming"}
	main()
	// Output:
	// 1, 2, 3, 4, 5, 6, 8, 9, 10, 12, 15, 16, 18, 20
}

func Example_main_Max5() {
	os.Args = strings.Fields("hamming -max 5")
	main()
	// Output:
	// 1, 2, 3, 4, 5
}

func Example_main_Max25() {
	os.Args = strings.Fields("hamming -max 25")
	main()
	// Output:
	// 1, 2, 3, 4, 5, 6, 8, 9, 10, 12, 15, 16, 18, 20, 24, 25
}

func Example_main_BadMaxOption() {
	os.Args = strings.Fields("hamming -max foo")
	main()
	// Output:
	// Could not convert max value to integer: strconv.Atoi: parsing "foo": invalid syntax
}

func Example_main_Max5_CommaSeparated() {
	os.Args = strings.Fields("hamming -max 5 -sep ,")
	main()
	// Output:
	// 1,2,3,4,5
}

func Example_main_Max5_NewlineSeparated() {
	os.Args = strings.Fields("hamming -max 5 -sep \\n")
	main()
	// Output:
	// 1
	// 2
	// 3
	// 4
	// 5
}

func Example_main_Max5_CommaTwoSpaceSeparated() {
	os.Args = []string{"hamming", "-max", "5", "-sep", ",  "}
	main()
	// Output:
	// 1,  2,  3,  4,  5
}

func Example_main_MaxOptionWithNoValue() {
	os.Args = []string{"hamming", "-max"}
	main()
	// Output:
}
