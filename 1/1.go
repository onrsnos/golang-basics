package main

// package greetings  // yeni bir package yaratıyor

import "fmt"

// Hello returns a greeting for the named person.
// func Hello(name string) string {
//     // Return a greeting that embeds the name in a message.
//     message := fmt.Sprintf("Hi, %v. Welcome!", name)
//     return message
// }

var packvar = "Package Scope"

func main() {

	// var name string  = "onur";

	degiskenString := "onur" // short decleration
	degiskenNumeric := 21
	degiskenBool := true

	var name = "bu bir stringdir :)"
	var uint8_degisken uint8 = 250    // 0-255
	var uint16_degisken uint16 = 4500 // 0-65536

	fmt.Println("String, ", degiskenString)
	fmt.Println("Numeric, ", degiskenNumeric)
	fmt.Println("Bool, ", degiskenBool)
	fmt.Println("name, ", name)

	fmt.Printf("%T\n", name)
	fmt.Printf("%T\n", degiskenNumeric)
	fmt.Printf("%T\n", degiskenBool)
	fmt.Printf("%T\n", name)

	fmt.Printf("%T\n", uint8_degisken)
	fmt.Printf("%T\n", uint16_degisken)

	var onur string // declaration

	var sen string = "Onur Sen" // assign (onur sen) // initial valyue = onur sen [ilk değer] // initialization [değişkeni başlatma]

	if true {
		var blockvar = "Block Scope"
		fmt.Println(blockvar)
	}

	var funcvar = "Function Scope"
	fmt.Println(funcvar)
	fmt.Println(packvar)

	fmt.Println(onur)
	fmt.Println(sen)

}
