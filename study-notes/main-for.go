package main

import (
	"fmt"
)

func main() {

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// the below similar to while loop since go doesn't have the word while

	j := 0
	for j <= 10 {
		fmt.Println(j)
		j++
	}

	sum := 0
	for {
		sum++
	}
	fmt.Println(sum) // unreachable code !!

	// for with continue
	for i := 0; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		} else {
			continue
		}
	}

	// for with break

	count := 0
	for i := 0; true; i++ {
		if i%13 == 0 {
			fmt.Printf("%d is dividable by 13\n", i)
			count++
		}
		if count == 10 {
			break
		}
	}
	fmt.Println("For loop has ended!")
}
