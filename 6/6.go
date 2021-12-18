package main

import (
	"fmt"
	"time"
	"bufio"
	"os"
)

func main() {

	fmt.Println("Solucan")

	// time

	fmt.Println(time.Now())

	reader := bufio.NewReader(os.Stdin) // a reader from the get a key from keyboard.

	value , err := reader.ReadString('\n')

	fmt.Println(value)
	fmt.Println(err)
}
