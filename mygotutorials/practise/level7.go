// Note: it is better to use slices in Go to store ordered set of values instead of an array
//Arrays:
package main

import "fmt"

func main() {

	var colors [3]string //size allocated
	colors[0] = "Red"
	colors[1] = "Green"
	colors[2] = "Blue" // values allocated
	fmt.Println(colors)
	fmt.Println(colors[0])

	var numbers = [5]int{1, 2, 3, 4, 5}
	fmt.Println(numbers)

	fmt.Println("Number of colors: ", len(colors))
	fmt.Println("Number of numbers: ", len(numbers))

}

// In Go we can use array only to store values. But to perform operations effectively on the ordered set of values we need to use slices
