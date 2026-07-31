package formas

import (
	"fmt"
	"math"
)

type Retangulo struct {
	altura  float64
	largura float64
}

type Circulo struct {
	raio float64
}

func (c Circulo) area() float64 {
	return math.Pi * math.Pow(c.raio, 2)
}

type Forma interface {
	area() float64
}

func EscreverArea(f Forma) {
	fmt.Printf("A forma é de %.02f", f.area())
	fmt.Println()

}

func (r Retangulo) area() float64 {
	return r.altura * r.largura
}

func main() {
	r := Retangulo{30, 10}
	c := Circulo{10}
	EscreverArea(c)
	EscreverArea(r)
}
