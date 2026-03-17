package main

import "fmt"

type Person struct {
	firstName string
	lastName  string
	age       int
	address   Address
}

type Address struct {
	city    string
	country string
}

func main() {

	p1 := Person{
		firstName: "John",
		lastName:  "Doe",
		age:       30,
		address: Address{ // embedded struct
			city:    "London",
			country: "U.K.",
		},
	}

	p2 := Person{
		firstName: "Peter",
		age:       23,
	}

	p2.address.city = "New York"
	p2.address.country = "USA"

	fmt.Println(p1.firstName)
	fmt.Println(p2.firstName)

	// Anonymous struct
	user := struct {
		userName string
		email    string
	}{
		userName: "user123",
		email:    "user@1243",
	}

	fmt.Println(user.email)

	fmt.Println(p1.fullName())
	fmt.Println(p1.address.city)
	fmt.Println(p2.address.country)

	fmt.Println("Before increment: ", p1.age)
	p1.incrementAgeByOne()
	fmt.Println("After increament: ", p1.age)

}

func (p Person) fullName() string {
	return p.firstName + " " + p.lastName
}

func (p *Person) incrementAgeByOne() {
	p.age++
}
