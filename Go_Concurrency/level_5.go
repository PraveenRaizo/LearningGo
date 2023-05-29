// the orders app - using Rest API  and Race Conditions in Go
/*
  Purpose:
     View all products and their amount of stock
	 Place a new order if it is valida and we have enough stock
	 View an existing order

	 GET/ - simply returns a hardcoded response
	 GET/products - returns all product information
	 POST/orders - place an order
	 GET/orders/{orderID} - get an placed order details using orderID
*/

/*
    - we will start 50 go routines in random order
	- each go routine sends the random order to the POST/orders endpoint using the net/http package (mutliple orders at same time)
	- we will use sync.WaitGroup to wait for the goroutines to complete
*/

/*
    Race Conditions: occurs when mutliple Go routines can read and write shared data without proper synchronization mechanisms
    In our case the go routines in this app will try to read and write data in a map. So we need to solve these Race conditions

	So we can use Go toolchain which can help us prevent race condition. add -race flag to any go command to use the toolchain.
	this will print stack traces and conflicting accesses when a race is found
*/

package main
