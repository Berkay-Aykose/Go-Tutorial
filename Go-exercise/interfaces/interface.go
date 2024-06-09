package interfaces

import (
	"fmt"
	"math"
)

type sekil interface {
	alan() float32
}

type Dikdortgen struct {
	width, height float32
}

type cember struct {
	yariCap float32
}

func (d Dikdortgen) alan() float32 {
	return d.width * d.height
}

func (c cember) alan() float32 {
	return math.Pi * c.yariCap * c.yariCap
}

func school(s sekil) {
	fmt.Println("Şeklin alanı : ", s.alan())
}

func Demo1() {
	x := cember{3}
	school(x)
	y := Dikdortgen{4, 5}
	school(y)

}
