package main

import "fmt"

var sayi = 25

func main() {
	switch true {
	case sayi < 25:
		fmt.Println("Sayı 25'den küçüktür")
		fallthrough
	case sayi < 50:
		fmt.Println("Sayı 50'den küçüktür")
		fallthrough
	case sayi < 75:
		fmt.Println("Sayı 75'den küçüktür")
		fallthrough
	case sayi < 100:
		fmt.Println("Sayı 100'den küçüktür")
	}
}
