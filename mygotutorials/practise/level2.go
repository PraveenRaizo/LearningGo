package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	// Let us see how to get input from a user

	input_reader := bufio.NewReader(os.Stdin) // reads an input from user
	fmt.Print("Enter text: ")
	// we use _ to indicate an variable that we are about to ignore/ unsure of. Here we use it to indicate a potential error in reading input
	// just like Java, Go too has - characters are wrapped in single quotes and full strings in double quotes
	input_received, _ := input_reader.ReadString('\n') // input_reader stops accepting input when user presses "enter" i.e: \n
	fmt.Println("You entered: ", input_received)

	// string type conversion. We are going to convert string into a different datatype (float64)
	fmt.Print("Enter a number: ")
	numInput, _ := input_reader.ReadString('\n')
	// here we care about the error as it is possible for the user to pass a string that can't be converted to integer so we use err word.
	aFloat, err := strconv.ParseFloat(strings.TrimSpace(numInput), 64) // 64 is the bit size float64. trimspace trims any unwanted space in user's input

	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Value of string input converted to float: ", aFloat) // output is a float64
	}

}
