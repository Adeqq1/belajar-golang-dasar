package main

import "fmt"

func main() {
	type noKtp string
	type married bool

	var noKtpAde noKtp = "12346531"
	var marriedStatus married = false

	fmt.Println(noKtpAde)
	fmt.Println(marriedStatus)

}
