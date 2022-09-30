package main

import "fmt"

func copyExample() {
	a6 := []int{-10, 1, 2, 3, 4, 5}
	a4 := []int{-1, -2, -3, -4}
	fmt.Println("a6:", a6)
	fmt.Println("a4:", a4)
	copy(a6, a4)
	fmt.Println("a6:", a6)
	fmt.Println("a4:", a4)
	fmt.Println()
}

func main() {
	s1 := make([]int, 5)
	reSlice := s1[1:3]
	fmt.Println(s1)
	fmt.Println(reSlice)

	reSlice[0] = -100
	reSlice[1] = 123456
	fmt.Println(s1)
	fmt.Println(reSlice)

	s2 := make([]int, 5)
	for i, _ := range s2 {
		s2 = append(s2, i)
	}

	fmt.Printf("Cap: %d, Length: %d\n", cap(s2), len(s2))

	s2 = append(s2, 100)
	fmt.Printf("Cap: %d, Length: %d\n", cap(s2), len(s2))
	copyExample()
}
