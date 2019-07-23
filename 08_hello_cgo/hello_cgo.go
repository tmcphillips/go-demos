package main

/*
#include <stdio.h>
#include <stdlib.h>
void write_message(const char * message) {
	printf("%s\n", message);
}
*/
import "C"
import "unsafe"

// Program that outputs a familiar greeting
func main() {
	message := C.CString("Hello World from CGO!")
	C.write_message(message)
	C.free(unsafe.Pointer(message))
}
