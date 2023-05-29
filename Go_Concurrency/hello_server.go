//using net/http for hello world

package main

import (
	"fmt"
	"net/http"
)

func main() {
	//configure the path and the handler function
	http.HandleFunc("/hello", Hello)
	//listen on port 8080 and block main
	fmt.Println("Listening on Port 8080...")
	http.ListenAndServe(":8080", nil)

}

func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, World!!!")
}

// In postman or browser try http://localhost:8080/hello after running this code
