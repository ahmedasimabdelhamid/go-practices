// Write a simple go app that does the following:
// 1. Print a total of an array (array will be hard coded in code)
// 2. Print Mean of an array
// 3. Sort array of strings alphabetically
// 4. Your functions should be inside two different packages (mymath package and mystring package)
// 5. Your main.go will import the functions and use them. Arrays will be hard coded in main function and passes as args.

package main

import (
	"fmt"

	"assigments/mystring"
)

func main() {
	sc1 := []string{"abc", "ssfe", "wqwe", "lol", "Ahh", "yalahwy"}
	mystring.StringAlphapitic(sc1)
	fmt.Println(sc1)
}
