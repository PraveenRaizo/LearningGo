//create a more advanced calculator app
package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func main() {
	value1 := getInput("Value 1")
	value2 := getInput("Value 2")

	var result float64

	switch operation := getOperation(); operation {
	case "+":
		result = addValues(value1, value2)
	case "-":
		result = subValues(value1, value2)
	case "*":
		result = mulValues(value1, value2)
	case "/":
		result = divValues(value1, value2)
	default:
		panic("Invalid Operator")
	}
	result = math.Round(result*100) / 100
	fmt.Printf("The result is %v\n\n", result)

}

func getInput(prompt string) float64 {
	fmt.Printf("%v ", prompt)
	input, _ := reader.ReadString('\n')
	value, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if err != nil {
		message := fmt.Sprintf("%v must be a number", prompt)
		panic(message)
	}
	return value

}

func getOperation() string {
	fmt.Println("Select an operation (+ - * /): ")
	op, _ := reader.ReadString('\n')
	return strings.TrimSpace(op)
}

func addValues(value1, value2 float64) float64 {
	return value1 + value2
}

func subValues(value1, value2 float64) float64 {
	return value1 - value2
}

func mulValues(value1, value2 float64) float64 {
	return value1 * value2
}

func divValues(value1, value2 float64) float64 {
	return value1 / value2
}
