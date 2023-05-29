package main

import "fmt"

const aConst int = 64 // this will be available for any function within this file

func main() {

	var hello string = "Hello!! Raizo!" // declaring a variable of type string
	var int1 int = 1
	hello1 := "You are a legend!!" // declaring a variable that takes the type of the value it is assigned with
	float1 := 2.2

	var anotherString = "Hello!! Gintoki!!" // declaring a variable that takes the type of the value it is assigned with

	// the := operator works only for variables inside a func. If outside func use var keyword

	fmt.Println(hello)
	fmt.Println(int1)
	fmt.Println(hello1) // has default new line feature
	fmt.Printf("The type of the variable hello is %T \n", hello)
	fmt.Printf("The type of the variable hello1 is %T \n", hello1) // printf returns the type of the variable but has no default new line. so use \n
	fmt.Printf("The type of the variable float1 is %T \n\n", float1)

	fmt.Println(anotherString)
	fmt.Printf("The type of the variable anotherString is %T\n\n", anotherString)

	fmt.Println("This is a const declared outside this function", aConst)
}
