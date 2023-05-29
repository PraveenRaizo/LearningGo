//Evaluate expressions with switch
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	rand.Seed(time.Now().Unix()) // calling seed function and passing time in UNIX format
	var dow int
	var result string

	switch dow = rand.Intn(7) + 1; dow { // dow is initalized with some random number
	case 1:
		result = "It's Sunday"
		//fallthrough             => if we use 'fallthrough' the after this case the flow will move to next case without break
	case 2:
		result = "It's Monday"
		//fallthrough  => if we use 'fallthrough' the after this case the flow will move to next case without break
	default:
		result = "It's someother day"

	}
	fmt.Println("Day: ", dow)
	fmt.Println(result)

}
