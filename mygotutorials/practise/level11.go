// program conditional logic
package main

import "fmt"

func main() {

	theAnswer := 42
	var result string

	if theAnswer < 0 { // in go we don't need () around the boolean
		result = "Less then zero"
	} else if theAnswer == 0 {
		result = "Equal to zero"
	} else {
		result = "Greater than zero"
	}

	fmt.Println(result)

	theAnswerNegative := -42

	negativeRes := checkGreater(theAnswerNegative)  //  calling another function and assigned it's returned value to negativeRes
	fmt.Println("Negative Result: ", negativeRes)

}

func checkGreater(y int) string {     // this function takes an integer parameter as input and returns string value
	
	var result string

	if x := y; x > 0 {     // we can initialize a variable in the conditional statement itself like we did with 'x' here
		result = "Greater than 0"
	} else if x == 0 {
		result = "Equal to 0"
	} else {
		result = "Less than 0"
	}

	return result
}