package loops

import (
	"fmt"
)

// asal sayı bulma
func Demo1() {
	asalsayi := 0
	fmt.Print("Bir sayı giriniz : ")
	fmt.Scanln(&asalsayi)
	asal := true

	for i := 2; i < asalsayi; i++ {

		if asalsayi%i == 0 {
			asal = false
		}

	}

	if !asal {
		fmt.Print("Asal sayı değildir..")
	} else {
		fmt.Print("Asal sayıdır...")
	}

}
