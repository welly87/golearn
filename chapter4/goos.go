package main

import(
	"fmt"
	"os"
)

func main() {
	var goos string = os.Getenv("GOOS")
	fmt.Println(goos)
	path := os.Getenv("PATH")
	fmt.Println(path)
}
