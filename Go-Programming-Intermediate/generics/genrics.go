package main

import "fmt"

type Stack[T any] struct {
	element []T
}

func swap[T any](a, b T) (T, T) {
	return b, a
}

func (s *Stack[T]) push(element T) {
	s.element = append(s.element, element)
}

func (s *Stack[T]) pop() (T, bool) {
	if len(s.element) == 0 {
		var zero T
		return zero, false
	}
	element := s.element[len(s.element)-1]
	s.element = s.element[:len(s.element)-1]

	return element, true
}

func (s *Stack[T]) isEmpty() bool {
	return len(s.element) == 0
}

func (s Stack[T]) printStack() {
	if len(s.element) == 0 {
		fmt.Println("Stack is empty!")
		return
	}

	fmt.Println("Stack elements: ")
	for _, ele := range s.element {
		fmt.Println(ele)
	}
}

func main() {
	x, y := 1, 2
	x, y = swap(x, y)
	fmt.Println(x, y)

	x1, y1 := "John", "Paul"
	x1, y1 = swap(x1, y1)
	fmt.Println(x1, y1)

	intStack := Stack[int]{}
	intStack.push(1)
	intStack.push(2)
	intStack.push(3)

	intStack.printStack()
	intStack.pop()
	intStack.printStack()
	fmt.Println("Is Stack empty: ", intStack.isEmpty())
}
