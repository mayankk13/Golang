package main

import "fmt"

func main() {
	day := "sunday"

	switch day {
	case "sunday":
		fmt.Println("It is weekend")
	case "monday":
		fmt.Println("It is weekday")
	case "tuesday":
		fmt.Println("It is weekday")
	default:
		fmt.Println("Invalid day!")
	}

	num := 2

	switch {
	case num > 1:
		fmt.Println("greater than 1")
		fallthrough
	case num == 2:
		fmt.Println("number is 2")
	default:
		fmt.Println("not 2")
	}
}
