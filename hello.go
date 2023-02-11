package main

import "fmt"

type User struct {
	name    string
	age     int
	address string
}

func main() {
	// Structs
	var u User
	u.address = "Tijuana, Baja california"
	u.age = 21
	u.name = "Francisco Rodríguez"
	fmt.Println(u)

	// Loops
	for i := 0; i < 5; i++ {
		fmt.Printf("Number %v\n", i+1)
	}

	// Functions
	fmt.Print(getName("Francisco Rodríguez"))

	// Arrays
	var array []int = make([]int, 100)
	array = append(array, 10)
	array[1] = 10
	fmt.Print(array)

	if u.age == 21 {
		fmt.Printf("\nAge: %v", u.age)
	}
}

func getName(n string) string {
	return "Name Function: " + n
}
