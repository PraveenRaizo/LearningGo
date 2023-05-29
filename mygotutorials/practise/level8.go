// Manage ordered value in slices
package main

import (
	"fmt"
	"sort"
)

func main() {

	var colors = [3]string{"red", "green", "blue"} //array
	fmt.Println("Colors array:", colors)

	var colorsSlice = []string{"red", "blue", "green", "purple"}
	fmt.Println("Colors slice:", colorsSlice)

	colorsSlice = append(colorsSlice, "orange") //appending new color to the slice and reassigning it to colorsSlice
	fmt.Println("Colors slice after append: ", colorsSlice)

	colorsSlice = append(colorsSlice[1:len(colorsSlice)]) //slices first element
	fmt.Println("After slicing first color: ", colorsSlice)

	colorsSlice = append(colorsSlice[:len(colorsSlice)-1]) //slices last element
	fmt.Println("After slicing last color: ", colorsSlice)

	numbers := make([]int, 5, 5) // using new() or make()
	numbers[0] = 130
	numbers[1] = 132
	numbers[2] = 131
	numbers[3] = 134
	numbers[4] = 133
	fmt.Println("Numbers Slice: ", numbers)

	numbers = append(numbers, 235)
	fmt.Println(numbers)

	sort.Ints(numbers) // using sort package to sort an slice. Use corresponding .
	fmt.Println("Numbers Slice after sort.Ints: ", numbers)

}
