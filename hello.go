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
}
