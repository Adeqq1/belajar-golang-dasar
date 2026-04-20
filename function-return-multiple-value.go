package main

import "fmt"

func getFullName() (string, string) {
	return "Ade", "Rifqy"
}

func main() {
	firstName, lastName := getFullName()
	fmt.Println(firstName, lastName)
}
