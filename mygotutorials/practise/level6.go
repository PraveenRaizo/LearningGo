// pointers:
package main

import "fmt"

func main() {

	anInt := 2
	var p = &anInt
	fmt.Println("value of p: ", *p)
	fmt.Println("Address of p: ", p)

	value1 := 42.13
	pointer1 := &value1
	fmt.Println("value 1: ", *pointer1)

	*pointer1 = *pointer1 / 31 // this changes the value of the variable value1 too
	fmt.Println("Pointer 1: ", *pointer1)
	fmt.Println("value1: ", value1)

}
