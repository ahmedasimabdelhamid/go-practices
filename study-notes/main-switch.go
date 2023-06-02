package main

import (
	"fmt"
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
	switch true {
	case n%2 == 0:
		fmt.Println("Even Number!")
	case n%2 != 0:
		fmt.Println("Odd Numner!")
	default:
		fmt.Println("Never Here!!")
	}

}
