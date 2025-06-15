package main

import (
	"fmt"
)

func main() {
	fmt.Print("Hello, World!\n")
	helloPrint("Nachiket")
	square, cube := returnvalue(5)
	fmt.Println("The square of 5 is:", square)
	fmt.Println("The cube of 5 is:", cube)
}

func helloPrint(name string) {
	fmt.Println("Hello, " + name + "!")
}

func returnvalue(a int) (square int, cube int) {
	square = a * a
	cube = a * a * a
	return
}
