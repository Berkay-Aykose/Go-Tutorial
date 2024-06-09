package structs

import "fmt"

func Struct() {
	fmt.Println(product{"Bilgisayar", 1, "xyz"})
	fmt.Println(product{"Bilgisayar", 1, ""})
}

type product struct {
	name  string
	price float32
	brand string
}
