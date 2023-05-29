//define functions as methods of custom types
package main

import "fmt"

func main() {

	poodle := Dog{"Poodle", 10, "Woof!!"}
	fmt.Println(poodle)
	fmt.Printf("%+v\n", poodle)
	fmt.Printf("Breed: %v\n Weight: %v\n", poodle.Breed, poodle.Weight)

	poodle.Speak()
}

//Dog is struct
type Dog struct {
	Breed  string
	Weight int
	Sound  string
}

func (d Dog) Speak() { // here (d, Dog) is a receiver. This is not function parameter. Also this is an exported function
	fmt.Println(d.Sound)
}

func (d Dog) SpeakThreeTimes() {
	d.Sound = fmt.Sprintf("%v %v %v", d.Sound, d.Sound, d.Sound)
	fmt.Println(d.Sound)
}
