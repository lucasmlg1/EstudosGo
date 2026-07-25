package main

import (
	"fmt"
	"math"
)

type retangulo struct {
	altura  float64
	largura float64
}

type circulo struct {
	raio float64
}

func (c circulo) area() float64 {
	return math.Pi * math.Pow(c.raio, 2)
}

type forma interface {
	area() float64
}

func escreverArea(f forma) {
	fmt.Printf("A forma é de %.02f", f.area())
	fmt.Println()

}

func (r retangulo) area() float64 {
	return r.altura * r.largura
}

func main() {
	r := retangulo{30, 10}
	c := circulo{10}
	escreverArea(c)
	escreverArea(r)
}
