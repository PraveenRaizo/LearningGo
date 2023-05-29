package main

import (
	"fmt"
	"math"
)

func main() {

	i1, i2, i3 := 12, 42, 75

	iSum := i1 + i2 + i3
	fmt.Println("Integers sum is: ", iSum)

	f1, f2, f3 := 23.5, 65.1, 76.3

	fSum := f1 + f2 + f3
	fmt.Println("Floats sum is : ", fSum)

	fSum = math.Round(fSum) // we are not using := becoz fSum is already initialized
	fmt.Println("Fully rounded float sum: ", fSum)

	fSum = f1 + f2 + f3
	fSum = math.Round(fSum*100) / 100 //mul and div by 100 for percision of 1 decimal point
	fmt.Println("Float sum with one decimal percision: ", fSum)

	circleRadius := 15.5
	circumference := circleRadius * math.Pi * 2
	fmt.Printf("Circumference of circle with radius %.2f: %.2f \n", circleRadius, circumference)

}
