package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	message := "Hello, \nGo!"
	message1 := "Hello \tGo!" // adds a tab space character
	message2 := "hello \rGo!" // takes cursor to the first character
	rawMessage := `Hello\nGo` // raw string, '\n' doesn't work

	fmt.Println(message)
	fmt.Println(message1)
	fmt.Println(message2)
	fmt.Println(rawMessage)

	// length
	fmt.Println("length of message variable is : ", len(message))
	fmt.Println("length of message1 variable is : ", len(message1))
	fmt.Println("length of message2 variable is : ", len(message2))
	fmt.Println("first char at message variable is : ", message[0])

	// concatination
	greeting := "hello "
	name := "Alice"
	fmt.Println(greeting + name)

	// comparison - lexico-graphical (comparison is based in ASCII value)
	str1 := "Apple"  // A has ASCII value - 65
	str := "apple"   // A has ASCII value - 97
	str2 := "banana" // b has ASCII value - 98
	str3 := "app"    // a has ASCII value - 97

	fmt.Println(str1 < str2)
	fmt.Println(str3 < str1)
	fmt.Println(str > str1)
	fmt.Println(str > str3)

	for i, char := range message {
		fmt.Printf("Character at index %d is %c \n", i, char)
		fmt.Printf("ASCII value of %c is %v\n", char, char)
	}

	fmt.Println("Rune count: ", utf8.RuneCountInString(message))
	fmt.Println("Rune count: ", utf8.RuneCountInString(greeting))

	// in Go - A rune is a alias for int32 and it represents a UniCode code point, a Unicode value
	// it is not a character, it is a integer value
	var ch rune = 'a'
	fmt.Println(ch)
}
