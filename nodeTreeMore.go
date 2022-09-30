package main

import "fmt"

func functioneOne(x int) {
	fmt.Println(x)
}

func main() {
	varOne := 1
	varTwo := 2
	fmt.Println("Hello There!")
	functioneOne(varOne)
	functioneOne(varTwo)
}
