package formas

import (
	"math"
	"testing"
)

func testArea(t *testing.T) {
	t.Run("Retângulo", func(t *testing.T) {
		ret := Retangulo{10, 12}
		areaEsperada := float64(120)
		areaRecebida := ret.area()
		if areaEsperada != areaRecebida {
			t.Errorf("A área recebida %f é diferente da área esperada %f", areaRecebida, areaEsperada)
		}
	})
	t.Run("Circulo", func(t *testing.T) {
		circle := Circulo{3}
		circuloRecebido := circle.area()
		circuloEsperado := float64(math.Pi * 9)

		if circuloEsperado != circuloRecebido {
			t.Errorf("O área do círculo %f é diferente da área esperada %f", circuloRecebido, circuloEsperado)
		}
	})
}
