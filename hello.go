package main

import "fmt"

type User struct {
	name    string
	age     int
	address string
}

func main() {
	var u User
	u.address = "Tijuana, Baja california"
	u.age = 21
	u.name = "Francisco Rodríguez"
	fmt.Println(u)

	for i := 0; i < 5; i++ {
		fmt.Printf("Number %v\n", i+1)
	}

	fmt.Print(getName("Francisco Rodríguez"))
}

func getName(n string) string {
	return "Name Function: " + n
}
