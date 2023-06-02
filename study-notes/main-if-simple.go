package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	i, err := strconv.Atoi("45")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(i)
	}

	if i, err := strconv.Atoi("25a"); err == nil {
		fmt.Println("No error!", i)
	} else {
		fmt.Println(err)
	}

	if args := os.Args; len(os.Args) != 2 {
		fmt.Println("One Argument is required!")
	} else if km, err := strconv.Atoi(os.Args[1]); err != nil {
		fmt.Println("the argument must be an intgeter! Error:", err)
	} else {
		fmt.Printf("%d km in miles is %v\n", km, float64(km)*1.609)
		_ = args
	}
}
