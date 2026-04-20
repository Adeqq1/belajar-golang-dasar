package main

import "fmt"

func main() {
	var (
		name1  = "Ade"
		name2  = "Ade"
		result = name1 == name2
	)

	fmt.Println(result)

	var (
		value1 = 100
		value2 = 200
	)

	fmt.Println(value1 == value2)
	fmt.Println(value1 > value2)
	fmt.Println(value1 < value2)
	fmt.Println(value1 != value2)

}
