package main

import "fmt"

func main() {
	var name string

	name = "Ade Rifqy Aulian"
	fmt.Println(name)

	name = "Ade Agung Kurniawan"
	fmt.Println(name)

	var daughtherName = "Aiza"
	fmt.Println(daughtherName)

	var age int8 = 20
	fmt.Println(age)

	country := "Indonesia"
	fmt.Println(country)

	var (
		firstName = "Ade"
		lastName  = "Aulian"
		fullName  = "Ade Aulian"
	)

	fmt.Println(firstName)
	fmt.Println(lastName)
	fmt.Println("Nama Panjang Saya Merupakan", fullName)
}
