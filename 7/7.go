package main

import (
	"fmt"
	// "time"
	"bufio"
	"os"
	"errors"
)

func main() {

	result, err := evenNum()
	// fmt.Println("Sonuç:", result)
	// fmt.Println(err)

	if err != nil {
		fmt.Println(err)
	} else{
		fmt.Println("Sonuç:", result)
	}

}

func evenNum() (string, error) {

	var num = "";
	reader := bufio.NewReader(os.Stdin) // a reader from the get a key from keyboard.

	num , err := reader.ReadString('\n')

	if err != nil {
		return "",errors.New("Hata: Bir sorun mevcut")
	} else {
		return num,nil
	}
}
