package main

import "fmt"

func main() {
	name := "Ade Rifqy Aulian"

	if name == "Ade Rifqy Aulian" {
		fmt.Println("Hello Ade Rifqy Aulian")
	} else if name == "Rifqy Aulian" {
		fmt.Println("Hello Rifqy Aulian")
	} else if name == "Eka" {
		fmt.Println("Hello Eka")
	} else {
		fmt.Println("Hello Stranger")
	}

	if length := len(name); length > 10 {
		fmt.Println("Nama Kamu Teralu Panjanng")
	} else {
		fmt.Println("Nama Kamu Sudah Benar")
	}
}
