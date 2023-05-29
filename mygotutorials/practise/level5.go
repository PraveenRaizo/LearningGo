// Simple Calculator
package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {

	fmt.Println("A simple calculator")

	input_reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter number num1: ")
	snum1, _ := input_reader.ReadString('\n')
	fnum1, err1 := strconv.ParseFloat(strings.TrimSpace(snum1), 64)
	if err1 != nil {
		fmt.Println(err1)
		panic("num1 input must be a number")
	}

	fmt.Println("Enter number num2: ")
	snum2, _ := input_reader.ReadString('\n')
	fnum2, err2 := strconv.ParseFloat(strings.TrimSpace(snum2), 64)
	if err2 != nil {
		fmt.Println(err2)
		panic("num2 input must be a number")
	}

	fsum := fnum1 + fnum2
	fsum = math.Round(fsum*100) / 100
	fmt.Printf("The sum of %v and %v is %v \n\n", fnum1, fnum2, fsum)

	if fnum1 > fnum2 {
		fsub := fnum1 - fnum2
		fmt.Printf("The diff between %v and %v is %v\n\n", fnum1, fnum2, fsub)
	} else {
		fsub := fnum2 - fnum1
		fmt.Printf("The diff between %v and %v is %v\n\n", fnum2, fnum1, fsub)
	}

	fmul := fnum1 * fnum2
	fmul = math.Round(fmul*100) / 100
	fmt.Printf("The product of %v and %v is %v\n\n", fnum1, fnum2, fmul)

}
