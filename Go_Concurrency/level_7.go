// ranging in channels
package main

import (
	"fmt"
	"time"
)

//greetings in many languages

var greetings = []string{"Hello!", "Ciao!", "Hola!", "Hej!", "Salut!"}

func main() {
	//create a channel
	ch := make(chan string, 1) // buffered channel
	//start a greeter to provide a greeting
	go greet(ch)
	//sleep for a long time
	time.Sleep(1 * time.Second)
	fmt.Println("Main Ready!!!")
	for greeting := range ch { // this channel closing method is better than the below commented 'ok' method
		//receive greeting
		// greeting, ok := <-ch
		// if !ok {
		// 	return // will close the main go routine if the channel is closed and no more greetings are coming
		// }
		//sleep and print
		time.Sleep(500 * time.Millisecond)
		fmt.Println("Greeting received!!!", greeting)
	}
}

// greet writes a greet to the given channel and then says goodbye
func greet(ch chan<- string) {
	fmt.Println("Greeter ready!!!")
	//greet
	for _, g := range greetings {
		ch <- g
	}
	close(ch) // without this the main goroutine's greeting := <-ch wouldn't know when to stop receiving and would throw error
	fmt.Println("Greeter completed!!!")
}
