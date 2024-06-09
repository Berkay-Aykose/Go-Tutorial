package maps

import "fmt"

func Map() {
	//key-value
	deger := make(map[string]int)
	//deger := map[string]string{...,}

	deger["Bir"] = 1
	deger["İki"] = 2
	deger["Üç"] = 3
	deger["Dört"] = 4
	deger["Beş"] = 5

	fmt.Println(deger["Bir"])
	fmt.Println("Adet :", len(deger))
	fmt.Println("Değer :", deger)
	delete(deger, "Bir")
	fmt.Println("Adet :", len(deger))
	fmt.Println("Değer :", deger)

}
