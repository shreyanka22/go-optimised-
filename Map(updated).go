package main

import (
    "fmt"
)

func main() {
    myMap := make(map[string]int)

    for {
        fmt.Println("\nMenu:")
        fmt.Println("1. Insert a key-value pair")
        fmt.Println("2. Update a value by key")
        fmt.Println("3. Delete a key-value pair")
        fmt.Println("4. Check existence of a key")
        fmt.Println("5. Print the map")
        fmt.Println("6. Exit")
        fmt.Print("Enter your choice (1-6): ")

        var choice int
        fmt.Scan(&choice)

        switch choice {
        case 1:
            fmt.Print("Enter key: ")
            var key string
            fmt.Scan(&key)
            fmt.Print("Enter value: ")
            var value int
            fmt.Scan(&value)
            myMap[key] = value
            fmt.Println("Key-value pair inserted.")
        case 2:
            fmt.Print("Enter key to update: ")
            var key string
            fmt.Scan(&key)
            fmt.Print("Enter new value: ")
            var value int
            fmt.Scan(&value)
            myMap[key] = value
            fmt.Println("Value updated.")
        case 3:
            fmt.Print("Enter key to delete: ")
            var key string
            fmt.Scan(&key)
            delete(myMap, key)
            fmt.Println("Key-value pair deleted.")
        case 4:
            fmt.Print("Enter key to check existence: ")
            var key string
            fmt.Scan(&key)
            value, exist := myMap[key]
            if exist {
                fmt.Printf("Value of %s: %d\n", key, value)
            } else {
                fmt.Printf("%s not found in map\n", key)
            }
        case 5:
            fmt.Println("Map contents:")
            for key, value := range myMap {
                fmt.Printf("Key: %s, Value: %d\n", key, value)
            }
        case 6:
            fmt.Println("Exiting the program.")
            return
        default:
            fmt.Println("Invalid choice. Please choose a valid option (1-6).")
        }
    }
}
