package main

import "fmt"

func main() {
	var a []int
	b := []int{}

	fmt.Println(a == nil)       // true
	fmt.Println(b == nil)       // false
	fmt.Println(len(a), cap(a)) // 0, 0
	fmt.Println(len(b), cap(b)) // 0, 0

	fmt.Println("--------------------------------------------------")

	c := []int{1, 2, 3}
	d := c[:2]

	d = append(d, 100)

	fmt.Println(c) // 1 2 3
	fmt.Println(d) // 1, 2

	fmt.Println("--------------------------------------------------")
}
