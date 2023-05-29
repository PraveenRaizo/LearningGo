// group related values in struct:
package main

import "fmt"

func main() {
	poodle := Dog{"Poodle", 10} //Dog object
	fmt.Println(poodle)
	fmt.Printf("%+v\n", poodle)
	fmt.Printf("Breed: %v\nWeight: %v\n", poodle.Breed, poodle.Weight)
	poodle.Weight = 9
	fmt.Printf("Breed: %v\nWeight: %v\n", poodle.Breed, poodle.Weight)
}

//Dog is a struct
type Dog struct {
	Breed  string
	Weight int
} // upper case means export in Go. Now Dog can be used in other parts of the application. If we declare this as dog it will throw error
