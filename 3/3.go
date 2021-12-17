package main

import "fmt"

func main() {

	F := -40

	K := float64(F-32)/1.8 + 273

	fmt.Printf("%v %T \n", K, K)

	deger := 27

	if deger < 0 {
		fmt.Println("Sıfırdan Küçük")
	} else if deger%2 == 0 {
		fmt.Println("Çift sayı")
	} else {
		fmt.Println("Tek sayı")
	}

	grade := 3 // 3/5

	switch grade {
	case 5:
		fmt.Println("Pekiyi")
	case 4:
		fmt.Println("İyi")
	case 3:
		fmt.Println("Orta")
	case 2:
		fmt.Println("Geçer")
	case 1:
		fmt.Println("Başarısız")
	}

	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			fmt.Println(i, " Çift Sayıdır.")

		} else {

			fmt.Println(i, " Tek Sayıdır.")
		}
	}
}
