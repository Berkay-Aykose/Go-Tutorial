package functions

//int döndürür
func Topla(sayi1 int, sayi2 int) int {
	var toplama = sayi1 + sayi2
	return toplama
}

func DortIslem(sayi1 int, sayi2 int) (int, int, int, float32) {
	Toplama := sayi1 + sayi2
	Cikarma := sayi1 - sayi2
	Carpma := sayi1 * sayi2
	Bolme := float32(sayi1 / sayi2)

	return Toplama, Cikarma, Carpma, Bolme
}
