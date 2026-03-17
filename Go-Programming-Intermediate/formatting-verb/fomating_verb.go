package main

import "fmt"

func main() {
	// General Formatting Verbs
	// %v	-	Prints the value in the default  format
	// %#v 	-	Prints the value in Go-Syntax format
	// %T	- 	Prints the type of the value
	// %%	- 	Prints the % sign

	i := 15.5
	string := "Hello World!"

	fmt.Printf("%v\n", i)
	fmt.Printf("%#v\n", i)
	fmt.Printf("%T\n", i)
	fmt.Printf("%v%%\n", i)

	fmt.Printf("%v\n", string)
	fmt.Printf("%#v\n", string)
	fmt.Printf("%T\n", string)

	num1 := 111_444
	num2 := 2

	fmt.Println(num1 + num2)

	// Integer formatting verbs
	// %b	-	Base 2
	// %d	-	Base10
	// %+d	-	Base 10 and always show sign
	// %o	- 	Base 8
	// %O	- 	Base 8, with leading 0o
	// %x	-	Base 16, lowercase
	// %X	-	Base 16, uppercase
}
