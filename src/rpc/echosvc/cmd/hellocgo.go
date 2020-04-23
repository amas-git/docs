package main

/*

#include <stdio.h>
const char* hello() {
	return "hello cgo";
}
*/
import "C"
import "fmt"

func main() {
	fmt.Println(C.GoString(C.hello()))
}
