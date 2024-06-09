package deferstatement

import "fmt"

func A() {
	fmt.Println("A fonsiyonu çalıştı...")
}

func B() {
	fmt.Println("B fonsiyonu çalıştı...")
}

func Demo1(yazi string) string {
	defer A()
	return yazi
	//defer B()  retundan sonrasını okumadığı için çalışmaz
}

func Test() {
	fmt.Println(Demo1("return olan kısım "))
}
