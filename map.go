package main

import "fmt"

func main() {
	person := map[string]string{
		"name":   "Ade Rifqy Aulian",
		"addres": "Talang Pantai",
	}

	person["title"] = "Programmer"
	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["addres"])

	book := make(map[string]string)
	book["title"] = "Belajar Golang"
	book["author"] = "Ade Rifqy Aulian"
	book["year"] = "2024"
	book["ups"] = "Ups"

	fmt.Println(book)
	fmt.Println(len(book))

	delete(book, "ups")

	fmt.Println(book)
	fmt.Println(len(book))
}
