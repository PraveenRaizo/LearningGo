// Channels: channels in action -  a simple hello program with channels
package main

import (
	"fmt"
	"time"
)

func main() {
	//create a channel
	ch := make(chan string) // unbuffered channel
	//ch := make(chan string, 1) //buffered channel
	//start the greeter to make a greeting
	go greet(ch)
	//sleep for a long time
	time.Sleep(5 * time.Second)
	fmt.Println("Main ready!!!")
	//receive greeting
	greeting := <-ch
	//sleep and print
	time.Sleep(2 * time.Second)
	fmt.Println("Greeting received!!!")
	fmt.Println(greeting)

}

func greet(ch chan string) {
	fmt.Printf("Greeter ready! \n Greeter waiting to send greeting....\n")
	//greet
	ch <- "Hello world"
	fmt.Println("Greeter completed!!!")
}

/* try to run the code using both unbuffered and buffered channels by commenting one of them while running another
   notice the difference
   Mostly the difference is greeter will be able to complete immediately even if main is not ready. This is becoz the buffered channel will
   be able to store the value, unlock the sender and keep the value until the main go routine is ready. The value is passed to main go routine when it's ready.
*/

/* NOTE:
   Bidirectional channel (able to support both send and receive) - chan T
   Send only channel  - chan<- T
   Receive only channel - <-chan T
   Allowed operations are enforced by compiler
   Bidirectional channels are implicitly cast to unidirectional channels but the reverse is not possible
*/
