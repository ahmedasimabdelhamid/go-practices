// slice like array but slice can shrink or grow
// slice , array , strings are slicible :: slicing doesn't modify the array or the slice it returns a new one!

package main

import (
	"fmt"
)

func main() {
	var cities []string
	fmt.Println("Cities is equal to nil: ", cities == nil)

	numbers := []int{1, 2, 3, 4, 5} // slice
	fmt.Println(numbers)

	nums := make([]int, 2)    // make is a built in function to create slice , 2 is the length of the slice
	fmt.Printf("%#v\n", nums) // result []int{0, 0}  slice of two zeros

	type names []string // define a new type called names and it's underlaying type is slice
	friends := names{"Dan", "Maria"}

	fmt.Println(friends)   // [Dan Maria]
	myfriend := friends[0] // accessing 1st element of slice
	fmt.Println(myfriend)

	friends[0] = "Gabrial" // modify 1st element
	fmt.Println(friends)
	for index, value := range friends {
		fmt.Printf("Index %v: value is %v \n", index, value)
	}

	var n []int           // not intialized
	fmt.Println(n == nil) // true

	m := []int{}          // empty but intialzed so it's not nil
	fmt.Println(m == nil) //  false

	a, b := []int{1, 2, 3}, []int{1, 2, 3} // two slices at same line

	// fmt.Println(a == b) // error a == b (slice can only be compared to nil)

	// to compare slices you have to iterate over them

	equal := true
	for index, valueOfA := range a {
		if valueOfA != b[index] {
			equal = false
		}
	}

	if len(a) != len(b) { // because if we set a = nil, it will not enter the for loop and the result will be identical

		equal = false
	}

	if equal {

		fmt.Println("a,b slices are identical")

	} else {

		fmt.Println("a,b slices are NOT identical !")

	}

	// to addend and copy slices

	newslice := []int{2, 3}
	newslice = append(newslice, 10) // append is not modiyfing the slice it returns a new one
	fmt.Println(newslice)
	newslice = append(newslice, 20, 30, 40)
	fmt.Println(newslice)

	newslice1 := []int{50, 60}
	newslice = append(newslice, newslice1...) // append newslice1 elements to existing newslice slice
	fmt.Println(newslice)

	// to copy a slice to another slice

	src := []int{10, 20, 30}
	dst := make([]int, 2)
	copyFromSrcToDst := copy(dst, src)

	fmt.Printf(" src is: %v\n dst is: %v\n number of elements copied is: %v\n", src, dst, copyFromSrcToDst)

	aa := [5]int{1, 2, 3, 4, 5}
	// a[start:stop]
	bb := aa[1:3]                   // [2 3]
	fmt.Printf("%v\n %T\n", bb, bb) // %T for type  (slicing an array returns a slice not array)
	fmt.Println(aa[1:])             // from 1st index to last index (to the length of the slice) called slice expression
	fmt.Println(aa[:3])             // from begining until 3rd element not included [1,2,3] called slice expression
	fmt.Println(aa[:])              // same as aa[0:len(aa)] called slice expression
	// fmt.Println(aa[:100])			// wil return a runtime error because the slice doesn't have 100 index
	bb = append(bb[:3], 100) // will append 100 after the 2nd element
	fmt.Println(bb)

	/* when creating a slice, behing the scenes go creates a hidden array called backing array
	    the backing array stores the elements not the slice, go implements the slice as data structure called
	   slice header.
	   	slice header contains 3 fields:
	   	- the address of the backing array (pointer).
	   	- the length of the slice len() returns it.
	   	- the capacity of the slice cap() returns it.
	   Notes :--------->
	   	- slice header is the runtime reprsentation of the slice.
	   	- A nil slice doesn't have a backing array.
	   	- a slice expression doesn't create a new backing array, the slice reurned by a slice expression refers
	   	  to the backing array of the orginal slice
	*/

	// below --> all three slices share the same backing array because we used slice expression to create the other two slices
	s2 := []int{1, 2, 3, 4, 5}
	s3, s4 := s2[0:2], s2[1:3]
	s3[1] = 600

	/* we only changed s3 however the other two slices have been modified and index 1 become 600 since they all
	share the same backing array due to slice expression */
	fmt.Println(s2) // [1 600 3 4 5]
	fmt.Println(s4) // [600 3]

	/* when we create a slice from array, the array become the backing array of the slice, below
	we changed the array value and the two slices has been modified same case as slice expression*/
	arr2 := [4]int{10, 20, 30, 40}
	s5, s6 := arr2[0:2], arr2[1:3]
	arr2[1] = 2
	fmt.Println(s5, s6) // [10 2] [2 30]

	// to create a complete new slice or a copy of an existing slice we should use append function
	cars := []string{"ford", "audi", "range"}
	newCars := []string{}

	// when appending elements of a slice to another slice don't forget the elipsis operator "..."
	newCars = append(newCars, cars[0:2]...) // newCars and cars don't share the same backing array they are completey different variables.

	cars[0] = "nissan"
	fmt.Println(cars, newCars) // newCarts didn't change!

	s7 := []int{10, 20, 30, 40, 50}
	newslice2 := s7[0:3]

	fmt.Println(len(s7), len(newslice2), cap(newslice2)) // cap represents number of elements of the backing array ---> 5 3 5
	newslice2 = s7[2:5]
	fmt.Println(newslice2)
	fmt.Println(len(s7), len(newslice2), cap(newslice2))

	var numss []int
	fmt.Printf("%#v\n", numss)
	fmt.Printf("Length: %d, Capacity: %d", len(numss), cap(numss))

}
