// create loops with for statements
package main

import "fmt"

func main() {
	colors := []string{"Red", "Green", "Blue"}
	fmt.Println(colors)

	for i := 0; i < len(colors); i++ {
		fmt.Println(colors[i])
	}

	fmt.Println("****************************")

	for i := range colors { // instead of len(colors) we can also use range
		fmt.Println(colors[i])
	}

	fmt.Println("****************************")

	for _, color := range colors { // here we are using the index to iterate, but we are ignoring it in the output or other part of the loop code. So we use _
		fmt.Println(color) // color
	}

	value := 1
	for value < 10 {
		fmt.Println("Value: ", value)
		value++
	}

	sum := 1
	for sum < 1000 {
		sum += sum
		fmt.Println("Sum : ", sum)
		if sum > 200 {
			goto theEnd
		}
	}

theEnd:
	fmt.Println("End of the program")
}
