package main

// #include <stdlib.h>
import "C"


func main() {
	Seed(9);
	println(Random())
}

func Random() int {
	return int(C.random())
}

func Seed(i int) {
	C.srandom(C.uint(i))
}
