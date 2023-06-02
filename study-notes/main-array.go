package main

import (
	"fmt"
	"strings"
)

func main() {
	var numbers [4]int // Array will be intialized by default with zeros
	fmt.Printf("%v\n", numbers)
	fmt.Printf("%#v\n", numbers) // Another way to print array [4]int{0, 0, 0, 0}

	// initialize array with given values
	// we can't add or remove from the array because its fixed length
	var array1 = [4]int{4, 5, 6, 7}
	array2 := [4]int{8, 9, 10, 11}
	array3 := [4]string{"Ahmed", "Asim"} // it's not mandatory to initialze the 4 valules, other two values are empty string
	array4 := [...]int{5, 3, 2, 4}       // ... caled elipsis means go will find out the array length you don't have to set it.
	array5 := [...]int{1,
		2,
		3,
		4,
		5, // this last comma is mandatory in case of multiline array
	}

	fmt.Printf("%#v\n%#v\n%#v\n%#v\n%#v\n", array1, array2, array3, array4, array5)
	fmt.Printf("the length of array4 is %d\n", len(array4))

	array1[0] = 7 // change array's first value
	// array1[5] = 8 //error because the array only has 3 elements
	fmt.Printf("%v\n", array1)

	// iterate over array :
	for index, value := range array1 {
		fmt.Println("index:", index, "value:", value)
	}

	// Adding ### 20 times

	fmt.Println(strings.Repeat("#", 20))
	// another way to iterate over array

	for i := 0; i < len(array1); i++ {
		fmt.Println(array1[i])
	}

	// multi daemntional array - array of arrays

	balances := [2][3]int{
		{5, 6, 7},
		[3]int{3, 5, 6},
	}

	fmt.Println(balances)

	// two arrays are equal if they have the same elements on the same order

	// if you copied the array to another array they are not connected and they stored in two different locations

	m := [3]int{1, 2, 4}
	n := m
	fmt.Println("m is equal to n", m == n) // will be evaluated to true
	m[0] = 2                               // change 1st index of m
	fmt.Println("m is equal to n", m == n) // will be evaluated to false

	grades := [3]int{
		1: 10,
		0: 5,
		2: 7,
	}
	accounts := [3]int{2: 50}      // only specify second element
	names := [...]string{5: "Dan"} // when I specified the 5th element the array will have 6(from 0 to 5) elements
	cities := [...]string{         // it will put pars at 4th and egypt at 3rd
		5: "London",
		"Paris",
		1: "Egypt",
	}
	fmt.Println(grades)
	fmt.Println(accounts)
	fmt.Printf("%#v\n", names)
	fmt.Printf("%#v", cities)

}
