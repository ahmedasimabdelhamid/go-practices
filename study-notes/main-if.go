package main
 
import (
    "fmt"
    // "strconv"
)
 
func main() {
 
    price, inStock := 100, true 
	if price > 80 {
		fmt.Println("Too Expensive!")
	}
	// _ = inStock

	if price <= 100 && inStock {
		fmt.Println("buy it!")
	}

	if price < 100 {
		fmt.Println("It's cheap!")
	} else if price == 100 {
		fmt.Println("On the edge!")
	} else {
		fmt.Println("Too expensive!")
	}

	age := 7
	if age >= 0 &&  age < 18 {
		fmt.Printf("You can not vote! Please return in %d years !\n", 18-age)
	}else if age == 18 {
		fmt.Printf("this is your first vote!")
	} else if age > 18 && age <= 100 {
		fmt.Printf("Please vote, it's importat")
	} else {
		fmt.Printf("Invalid age!")
	}
}