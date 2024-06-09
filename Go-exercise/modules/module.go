package modules

import "fmt"

// fonksiyonları büyük harfle başlat gerekir.
func Modul() {
	fmt.Println("merhaba Dünya")

	var metin string = "berkay"
	fmt.Println(metin)

	var sayi int = 10
	fmt.Println(sayi)

	var sayi2 float64 = 0.5
	fmt.Println(sayi2)

	sayi3 := 20
	fmt.Println(sayi3)
	fmt.Printf("Veri tipi : %T\n", sayi3)

	var durum bool = false

	durum = (sayi == sayi3)
	fmt.Println(durum)
}
