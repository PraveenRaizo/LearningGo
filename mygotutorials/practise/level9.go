// store unordered values in maps
package main

import (
	"fmt"
	"sort"
)

func main() {

	secondName := make(map[string]string) // a string to string mapping and using make() to initialize a dynamic size
	fmt.Println(secondName)

	secondName["Naruto"] = "Uzumaki"
	secondName["Sasuke"] = "Uchiha"
	secondName["Sakura"] = "Haruno"

	fmt.Println("Map with values: ", secondName)

	narutoLastName := secondName["Naruto"]          //using key to get the value in map and assign it to a var
	fmt.Println("narutoLastName: ", narutoLastName) // result must be Uzumaki

	delete(secondName, "Sakura")
	fmt.Println("Map after deleting Sakura data:", secondName)

	secondName["Kakashi"] = "Hatake"
	fmt.Println("Map after adding Kakashi data: ", secondName) // however this adds Kakashi as the first data in the map as the map is ordered based on ascending order of first letter of the key

	fmt.Println("\n\n**********\n\nPrinting the Map using for loop: ")
	for key, value := range secondName {
		fmt.Printf("%v : %v \n", key, value)
	}

	// for custom ordering of the Map in alphabetical order
	// in our case it is already ordered in alphabetical order
	keys := make([]string, len(secondName))
	i := 0
	for k := range secondName {
		keys[i] = k
		i++
	}
	fmt.Println(keys)
	sort.Strings(keys)
	fmt.Println(keys)
}
