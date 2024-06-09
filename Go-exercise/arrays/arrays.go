package arrays

import "fmt"

func Dizi() {

	var sayilar [5]int
	//[ 0 0 0 0 0]
	sayilar[2] = 50
	//[ 0 0 50 0 0]

	for i := 0; i < 5; i++ {
		fmt.Println(sayilar[i])
	}

	demoSayilar := [5]int{20, 30, 40, 50, 60}

	for i := 0; i < len(demoSayilar); i++ {
		fmt.Println(demoSayilar[i])
	}
}
