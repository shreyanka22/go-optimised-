package main

import "fmt"

func main() {
	myMap := map[string]int{
		"apple":  1,
		"banana": 2,
		"orange": 3,
	}

	appleValue := myMap["apple"]
	bananaValue := myMap["banana"]
	fmt.Println("Value of apple :", appleValue)
	fmt.Println("Value of Banana :", bananaValue)

	myMap["apple"] = 5

	fmt.Println("Updated value of apple :", myMap["apple"])

	delete(myMap, "orange")
	fmt.Println("After deleting orange :", myMap)

	if value, exists := myMap["banana"]; exists {
		fmt.Println("Value of banana :", value)
	} else {
		fmt.Println("Banana not found in the map.")
	}

	for key, value := range myMap {
		fmt.Println("Key :", key, "Value :", value)
	}

	fmt.Println("Length of the map :", len(myMap))
}
