package main

import (
	"fmt"
)

func calcArea(r float64) float64 {
	return 3.14 * r * r
}

func calcPerimeter(r float64) float64 {
	return 2 * 3.14 * r
}

func calcDiameter(r float64) float64 {
	return 2 * r
}

func printResult(radius float64, calcFuction func(r float64) float64) {
	result := calcFuction(radius)
	fmt.Println("Results: ", result)
	fmt.Println("Thank You")
}


func getFunction(query int) func(r float64) float64 {
              query_to_func := map[int]func(r float64) float64{
				1: calcArea,
				2: calcPerimeter,
				3: calcDiameter,
			  }
			  return query_to_func[query]
}


func main() {
	var query int
	var radius float64
	fmt.Print("Enter the raduis of the circle: ")
	fmt.Scanf("%f", &radius)
	fmt.Printf("Enter \n 1 - area \n 2 - perimeter \n 3 - diameter: ")
	fmt.Scanf("%d", &query)

printResult(radius, getFunction(query))

	

}
