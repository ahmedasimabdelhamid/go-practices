package main
import (
	"fmt"
)


func main() {
	people := [5]string {"helen", "mark", "brenda", "mary", "micheal"}
	friends := [2]string {"mark", "mary"}

	outer: // definging a label called outer 
	  for index, name := range people {   // range iterats over array and return both index and value
		 for _, friend := range friends {
			if name == friend { 
				fmt.Printf("Found a friend %q at index %d\n", friend, index )
				break outer
			}
		 }
	  }
	  fmt.Printf("Next instruction after the break!") 

}