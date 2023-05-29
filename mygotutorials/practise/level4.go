//time and date:
package main

import (
	"fmt"
	"time"
)

func main() {

	n := time.Now()
	fmt.Println("Current time: ", n.Format(time.ANSIC)) //formatting time to ANSIC format for easy readability

	t := time.Date(2023, time.May, 26, 14, 0, 0, 0, time.UTC) // setting a time ourself in UTC format
	fmt.Println("Tomorrow afternoon: ", t.Format(time.ANSIC)) // formatting it to ANSIC format for easy readability

	parsedTime, _ := time.Parse(time.ANSIC, "Tue Nov 10 23:00:00 2009") // getting back the time obj from ANSIC time
	fmt.Printf("The type of the parsedTime is %T\n", parsedTime)

}
