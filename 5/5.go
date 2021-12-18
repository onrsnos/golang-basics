package main

import "fmt"

const name = "Onur"
const surName = "Şen"

func main() {
	fmt.Println("Deneme")

	var x, y, sum, dec int

	// x = 5
	// y = 10

	// sum = x + y

	// fmt.Println("Toplam: ", sum)

	// x = 5
	// y = 10

	// dec = y - x

	// fmt.Println("Fark: ", dec)

	// // Moduler Programming
	// // Readlable Code
	// // From complex to Simple

	x = 5
	y = 10

	sum = sumFunction(x, y)
	fmt.Println("Toplam: ", sum)

	x = 5
	y = 10

	dec = decFunction(y, x)
	fmt.Println("Fark: ", dec)

	getNames()
}

func sumFunction(x, y int) int {
	return x + y
}

func decFunction(x, y int) int {
	return x - y
}

func getNames() {
	fmt.Println("Merhaba ben", name, surName)
}
