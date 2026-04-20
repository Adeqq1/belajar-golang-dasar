package main

import "fmt"

func main() {
	name := "Ade"

	switch name {
	case "Ade":
		fmt.Println("Halo Ade")
	case "Eko":
		fmt.Println("Halo Eko")
	default:
		fmt.Println("Halo Stranger")
	}

	// switch length := len(name); length > 5 {
	// case true:
	// 	fmt.Println("Nama terlalu panjang")
	// case false:
	// 	fmt.Println("Nama sudah benar")
	// }

	length := len(name)
	switch {
	case length > 10:
		fmt.Println("Nama terlau panjang")
	case length > 5:
		fmt.Println("Nama lumayan panjang")
	default:
		fmt.Println("Nama sudah benar")
	}
}
