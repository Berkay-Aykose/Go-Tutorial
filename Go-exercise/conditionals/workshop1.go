package conditionals

import (
	"fmt"
	"math/rand"
)

func Demo1() {
	//üç adet int değişken belirleyip
	//büyük olanı ekrana yazdırınız

	sayi_1 := rand.Intn(100) + 1
	fmt.Println("1. sayı:", sayi_1)

	sayi_2 := rand.Intn(100) + 1
	fmt.Println("2. sayı:", sayi_2)

	sayi_3 := rand.Intn(100) + 1
	fmt.Println("3. sayı:", sayi_3)

	var enBuyuk = sayi_1

	if enBuyuk < sayi_2 {
		enBuyuk = sayi_2
	} else if enBuyuk < sayi_3 {
		enBuyuk = sayi_3
	}

	fmt.Printf("en büyük sayı %v", enBuyuk)
}
