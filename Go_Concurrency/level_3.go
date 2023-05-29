// in previous level we saw using sleep was not optimal way for concurrency
// in this level we will use sync package which is the proper way to handle concurrency

/*
  We will use sync.WaitGroup:
    - Used to wait for goroutines to finish
	- under the hood, it uses a very simple counter and inner lock
	- the zero value of the wait group is ready to be used

	Methods of sync.WaitGroup:
	- func (wg *WaitGroup) Add(delta int) - adds a given number to the inner counter. This could be the number of goroutines we wish to wait for. Add also panics if the intercounter becomes negative
	- func (wg *WaitGroup) Done() - decrements the inner counter by 1. Should be used when a goroutine finishes its work
	- func (wg *WaitGroup) Wait() - wait blocks the goroutine from which it is invoked. We can use this to replace the sleep.time()
*/
package main

import (
	"fmt"
	"sync"
)

func main() {

	var wg sync.WaitGroup
	wg.Add(1)     // we are going to wait for only 1 go routine: hello goroutine
	go hello(&wg) // we are passing wg into hello
	wg.Wait()     // we replaced this: time.Sleep(1 * time.Second) will wait until the innercounter of main goroutine goes to 0
	goodbye()
}

func hello(wg *sync.WaitGroup) {
	defer wg.Done() // indicates that the hello goroutine is done its work
	fmt.Println("Hello World!!!!")
}

func goodbye() {
	fmt.Println("Goodbye!!! Fools!!!")
}

/* NOTE!!!!!!!:
     Without &wg and *wg sync.WaitGroup we will get this error:
			fatal error: all goroutines are asleep - deadlock!

		goroutine 1 [semacquire]:
		sync.runtime_Semacquire(0xc0000bc000?)
				/usr/local/go/src/runtime/sema.go:56 +0x25
		sync.(*WaitGroup).Wait(0x0?)
				/usr/local/go/src/sync/waitgroup.go:136 +0x52
		main.main()
				/home/agira/Desktop/Go lang Tutorials/LearningGo-master/Go_Concurrency/level_3.go:27 +0x7d
		exit status 2


		This is because of the way Go passes parameters to its functions. Though we use wg.Add(1) and initialise the counter to 1, this only turns the inner counter of main goroutine from 0 to 1
		However at go hello(wg) hello goroutine is initialised with counter 0. When we use defer wg.Done() this counter goes to -1 and throws error
		This occurs because go passes params by value to functions. As a result under the hood a copy waitgroup with counter 0 is initialized

		To fix this we pass param by reference instead of by value. Hence the &wg and *wg sync.WaitGroup

*/
