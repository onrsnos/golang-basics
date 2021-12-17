package main

import "fmt"

func main() {

	F := -40

	K := float64(F-32)/1.8 + 273

	fmt.Printf("%v %T \n", K, K)
}
