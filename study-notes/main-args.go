package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	fmt.Println("os.Args", os.Args)
	fmt.Println("Path:", os.Args[0])
	fmt.Println("firstArgument:", os.Args[1])
	fmt.Println("secondArgument:", os.Args[2])
	fmt.Println("thirdArgument:", os.Args[3])
	fmt.Println("Number of Arguments inside os.Args is:", len(os.Args))

	var result, err = strconv.ParseFloat(os.Args[1], 64)
	fmt.Printf("%T\n", result)
	_ = err
}
