package main

import (
	"fmt"
)

type Weekday int

const (
	Sunday Weekday = iota
	Monday
	Tuesday
	Wednesdday
	Thursday
	Friday
	Saturday
)

func main() {
	fmt.Println("day of the week ", Tuesday)
}

func (d Weekday) String() string {
	return [...]string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}[d]
}
