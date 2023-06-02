package main

import (
	"fmt"
	"time"
)

func main() {
	// switch is same as if statement but the purpose of switch is to make very long if statement more readable

	language := "Golang"
	switch language {
	case "python":
		fmt.Println("you are learning python")
	case "Go", "golang":
		fmt.Println("Good, Go for go!")
	default:
		fmt.Println("Any other programming language is a good start")
	}

	n := 5
	switch true { // same as switch {
	case n%2 == 0:
		fmt.Println("Even Number!")
	case n%2 != 0:
		fmt.Println("Odd Numner!")
	default:
		fmt.Println("Never Here!!")
	}

	// print time now
	hour := time.Now().Hour()
	fmt.Println(hour)

	switch { // means true

	case hour < 12:
		fmt.Println("Good morning!")
	case hour < 17:
		fmt.Println("Good Afternoon!")
	default:
		fmt.Println("Good Evening!")
	}
}
