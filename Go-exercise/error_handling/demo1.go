package errorhandling

import (
	"fmt"
	"os"
)

func Demo1() {
	// 'os' go nun bize sağladığı bir kütüphane
	//burada deneme.txt dosyasını açmaya yarıyor
	f, err := os.Open("error_handling/deneme.txt") //isim değiştirilse öyle bir dosya olmadığı için hata verecek
	//nil = null

	if err != nil {
		if pErr, ok := err.(*os.PathError); ok {
			fmt.Println("Dosya bulunamadı : ", pErr.Path) //yanlış aradığın dosya adını yazdırır
			return
		} else {
			fmt.Println("Dosya bulunamadı")
			return
		}

	} else {
		fmt.Println(f.Name())
	}
}
