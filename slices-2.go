package main

import "fmt"

func main() {
	aSlice := []int{1, 2, 3, 4, 5}
	fmt.Println(aSlice)
	integers := make([]int, 2)
	fmt.Println(integers)
	integers = nil
	fmt.Println(integers)
	anArray := [5]int{-1, -2, -3, -4, -5}
	refAnArray := anArray[:]
	fmt.Println(anArray)
	fmt.Println(refAnArray)
	anArray[4] = 100
	fmt.Println(refAnArray)
	// One dimensional slice
	single := make([]byte, 5)
	fmt.Println(single)
	// two dimensions
	double := make([][]int, 3)
	fmt.Println(double)
	fmt.Println()

	for i := 0; i < len(double); i++ {
		for j := 0; j < 2; j++ {
			double[i] = append(double[i], i*j)
		}
	}

	for _, single := range double {
		for i, y := range single {
			fmt.Println("i:", i, "value:", y)
		}
		fmt.Println()
	}
}
