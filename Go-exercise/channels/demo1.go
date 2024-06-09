package channels

import (
	"fmt"
	"time"
)

func CiftSayilar(ciftSayiCn chan int) {
	top := 0
	for i := 0; i <= 10; i = i + 2 {
		fmt.Println("Cift sayi : ", i)
		top = top + i
		time.Sleep(1 * time.Second)
	}
	ciftSayiCn <- top
}

func TekSayilar(tekSayiCn chan int) {
	top := 0
	for i := 1; i <= 10; i = i + 2 {
		fmt.Println("Tek sayi : ", i)
		top = top + i
	}
	tekSayiCn <- top
}

func Demo1() {

}
