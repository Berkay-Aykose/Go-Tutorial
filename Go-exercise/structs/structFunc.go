package structs

import "fmt"

type customer struct {
	firsName string
	lastName string
	age      int
}

func (c customer) save() {
	fmt.Println("Eklendi : ", c.firsName)
	fmt.Println("Eklendi : ", c.age)
}

func Demo1() {
	c := customer{firsName: "Berkay", lastName: "Ayköse", age: 20}
	c.save()
}
