// using go keyword we can start a new goroutine

/*
  In a go program what happens under the hood:
    1. Allocate memory for the program
	2. Start a main goroutine for the program
	3. Run the program on a thread
	4. Execute the code
	5. Shut down and clean up once the program completes
*/

package main

import (
	"fmt"
	"time"
)

// this first main is without timer.

/*

func main() {

	go hello()
	goodbye()
}

func hello() {
	fmt.Println("Hello World!!!!")
}

func goodbye() {
	fmt.Println("Goodbye!!! Fools!!!")
}
*/

/*
  If you try running the above code you will get Goodbye printed always, but Hello will be always printed after Goodbye sometimes and somtimes it won't be printed at all
  Reason: suppose at time t1 the main goroutine starts.
          At time t2 it invokes the hello goroutine and the hello function
		  Starting hello goroutine doesn't block the main goroutine. So concurrently the main goroutine runs and calls the goodbye function
		  So Goodbye gets printed always and once Goodbye is printed the main goroutine ends.
		  When main goroutine ends all its associated resources including hello goroutine ends.
		  Sometimes the hello goroutine manages to print Hello World before main goroutine ends and sometimes it doesn't.
		  This is why hello world is printed sometimes and not printed sometimes.


		  To prevent this we can add a time delay to stop the main goroutine from ending before the hello goroutine is finished.
*/

func main() {
	go hello()
	time.Sleep(1 * time.Second)
	goodbye()
}

func hello() {
	fmt.Println("Hello World!!!!")
}

func goodbye() {
	fmt.Println("Goodbye!!! Fools!!!")
}

// this time using sleep we are pausing the main goroutine which allows hello goroutine will print Hello World
// however sleep doesn't guarantee an order in results. So sleep must not be used as a production level solution
