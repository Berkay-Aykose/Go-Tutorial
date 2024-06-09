package conditionals

import "fmt"

func If() {
	var hesap float32 = 1000
	var cekilmekIstenen float32 = 900

	if hesap < cekilmekIstenen {
		fmt.Println("Yeterli para yok")
	} else {
		fmt.Println("Para çekildi")
		//hesap = hesap -cekilmekIstenen
		hesap -= cekilmekIstenen
	}

	fmt.Printf("Hesaptaki kalan para : %v", hesap)

}
