package errorhandling

import (
	"fmt"
)

func dogrula(i interface{}) {
	sayi, ok := i.(int) //sayının int olup olmadını kontrol eder "ok" ona göre true yada false gönderirSW
	fmt.Println(sayi, ok)
}
func Demo2() {
	//dogrula(3) aşşadaki gibide tanımlanabilir

	var x interface{} = 3
	dogrula(x)

	var y interface{} = "Berkay"
	dogrula(y)
}
