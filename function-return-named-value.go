package main

import "fmt"

func getFullname2() (firtsName, middleName, lastName string) {
	firtsName = "Ade"
	middleName = "Rifqy"
	lastName = "Aulian"

	return
}

func main() {
	a, b, c := getFullname2()
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(a, b, c)
	fmt.Println(getFullname2())
}
