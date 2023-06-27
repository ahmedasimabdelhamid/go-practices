package main

import (
	"fmt"
	"strings"
)

func main() {

	// an anonymous struct is a struct with no explicitly defined struct type alias.
	diana := struct {
		firstName, lastName string
		age                 int
	}{
		firstName: "Diana",
		lastName:  "Muller",
		age:       30,
	}

	fmt.Printf("%#v\n", diana)
	// =>struct { firstName string; lastName string; age int }{firstName:"Diana", lastName:"Muller", age:30

	//** ANONYMOUS FIELDS **//

	// fields type becomes fields name.
	type Book struct {
		string
		float64
		bool
	}

	b1 := Book{"1984 by George Orwell", 10.2, false}
	fmt.Printf("%#v\n", b1) // => main.Book{string:"1984 by George Orwell", float64:10.2, bool:false}

	fmt.Println(b1.string) // => 1984 by George Orwell

	// mixing anonymous with named fields:
	type Employee1 struct {
		name   string
		salary int
		bool
	}

	e := Employee1{"John", 40000, false}
	fmt.Printf("%#v\n", e) // => main.Employee1{name:"John", salary:40000, bool:false}
	e.bool = true          // changing a field

	fmt.Println(strings.Repeat("#", 10))

	//** EMBEDDED STRUCTS **//
	// An embedded struct is just a struct that acts like a field inside another struct.

	// define a new struct type
	type Contact struct {
		email, address string
		phone          int
	}

	// define a struct type that contains another struct as a field
	type Employee struct {
		name        string
		salary      int
		contactInfo Contact
	}

	// declaring a value of type Employee
	john := Employee{
		name:   "John Keller",
		salary: 3000,
		contactInfo: Contact{
			email:   "jkeller@company.com",
			address: "Street 20, London",
			phone:   042324234,
		},
	}

	fmt.Printf("%+v\n", john)
	// => {name:John Keller salary:3000 contactInfo:{email:jkeller@company.com address:Street 20, London phone:295619381404}}

	// accessing a field
	fmt.Printf("Employee's salary: %d\n", john.salary)

	// accessing a field from the embedded struct
	fmt.Printf("Employee's email:%s\n", john.contactInfo.email) // => Employee's email:jkeller@company.com

	// updating a field
	john.contactInfo.email = "new_email@thecompany.com"
	fmt.Printf("Employee's new email address:%s\n", john.contactInfo.email)
	// =>  Employee's new email address:new_email@thecompany.com
}
