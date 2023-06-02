package main
import (
	"fmt"
)
func main() { 
	

	// it creates a for loop exact like for does
i := 0
loop: 
  if i < 5 {
	fmt.Println(i)
	i++
	goto loop 
  }

// the below will not work because goto can't work with variable declared below it
goto todo
x := 5
todo: 
 fmt.Println("something here")
}

