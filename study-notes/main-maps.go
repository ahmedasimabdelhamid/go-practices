package main

import (
	"fmt"
)

func main() {

	var employees map[string]string
	fmt.Printf("%#V\n", employees)
	fmt.Printf("No of pairs: %d\n", len(employees))                     // length of map
	fmt.Printf("The value for key %q is %q\n", "Dan", employees["Dan"]) // get value of key in a map

	var accounts map[string]float64
	fmt.Printf("%#v\n", accounts["0x323"])

	/*
		we can't use slices as a key in a map because slices are not comparable we can't compare slices
		but we can use arrays as keys
	*/

	// all keys in a map must have the same type, all values must have the same values
	people := map[string]float64{}
	people["John"] = 21.4
	people["Marry"] = 10
	fmt.Println(people)

	balances := map[string]float64{
		"USD": 323.11,
		"EUR": 432.13,
		// 50: 322.1,    //Error, all keys must bthe same type
		"CHF": 3243.1, // last comma is manadatory in multiline value
	}

	m := map[int]int{1: 10, 2: 20, 3: 30}
	_ = m
	fmt.Println(balances)

	balances["USD"] = 4344.112       // updating the value of USD
	balances["SAR"] = 12203121132133 // adding new key value
	fmt.Println(balances)

	// looking for non existing key in a map will retrun zero
	// sometimes you need to distinguish between the zero value or the non exiting value so we use comma ok idiom

	fmt.Println(balances["RON"])
	v, ok := balances["RON"] // the first v is the valuem the second ok is a bool if exist will be true if not will be false and that is how we distinguish

	if ok {
		fmt.Println("the RON balance is: ", v)
	} else {
		fmt.Println("The RON key doesn't exist in the map!")
	}

	// maps in go is designed for fast lookup time not looping

	for k, v := range balances {
		fmt.Printf("Key: %#v, Value: %v\n", k, v)
	}

	delete(balances, "USD")
	fmt.Println(balances)

	a := map[string]string{"A": "X"}
	b := map[string]string{"B": "Y"}

	// fmt.Println(a == b)  maps can't be compared, maps can only compared to nil
	// this is one way to compare maps where the keys and values are of type string
	s1 := fmt.Sprintf("%s", a) // get the representative string of a map a
	s2 := fmt.Sprintf("%s", b) // get the representative string of a map b

	fmt.Println(s1, s2)

	if s1 == s2 {
		fmt.Println("Maps are equal")
	} else {
		fmt.Println("Maps are not equal")
	}

	/* when declaring a map go creates a pointer to a map header in memory
	   the map references this internal data structure the map header
	   the map contains only the memory address of the map header, the key value pairs of the map
	   are not stored directly into the map, there are stored in memory and the adddress referenced by the map header
	   when you copy a map to a new map, the internal data structure is not copied, just referenced
	   both maps have the same map header. */

	friends := map[string]int{"Dan": 40, "maria": 25}
	neighbors := friends // neighbors is not a copy of friends, it referneces the same data structure in memory

	friends["Dan"] = 50
	fmt.Println(neighbors, "\n") // value of Dan is 50

	// to copy a map to a new map we intialize the new map then for loop over the old and copy to the new

	people1 := make(map[string]int)

	for k, v := range friends {
		people1[k] = v
	}

	people1["Anne"] = 18 // when we modified people1, friens stays the same
	fmt.Printf("%#v ------------------%#v\n", people1, friends)

	delete(friends, "Dan") // delete from friends
	fmt.Println(friends)

	delete(people, "andire") // andire doesn't exist there is no error returned
}
