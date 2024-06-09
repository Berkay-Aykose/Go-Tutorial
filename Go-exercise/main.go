package main

import (
	"fmt"
	"golesson/restful"
)

func main() {
	fmt.Println("merhaba Dünya")
	//modules.Modul()
	//conditionals.If()
	//conditionals.Demo1()
	//loops.For()
	//loops.Demo1()
	//arrays.Dizi()
	//slices.Slice()
	//fmt.Println(functions.Topla(2, 3))

	/*
		sonuc1, sonuc2, sonuc3, sonuc4 := functions.DortIslem(10, 2)
		fmt.Println("Toplama :", sonuc1)
		fmt.Println("Çıkarma :", sonuc2)
		fmt.Println("Çarpma :", sonuc3)
		fmt.Println("Bölme :", sonuc4)
	*/

	//fmt.Println(functions.ToplaVariadic(5, 3, 7, 6))

	//maps.Map()
	//for_range.Range()
	//structs.Demo1()

	/*
		ciftSayiCn := make(chan int)
		tekSayiCn := make(chan int)
		go channels.CiftSayilar(ciftSayiCn)
		go channels.TekSayilar(tekSayiCn)

		ciftSayiToplam, tekSayiToplam := <-ciftSayiCn, tekSayiCn

		carpim := ciftSayiToplam * <-tekSayiToplam

		fmt.Println("Çarpım : ", carpim)
	*/

	//interfaces.Demo1()

	//deferstatement.Test()

	//errorhandling.Demo1()
	//errorhandling.Demo2()
	//fmt.Println(errorhandling.TahminEt(333))

	restful.Demo2()
}
