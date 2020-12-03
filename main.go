package goArea

import "math"

// Pi é sempre o mesmo
const Pi = 3.1416

//Circ vai calcular area de um circulo
func Circ(raio float64) float64 {
	return math.Pow(raio, 2) * Pi
}

//Rect vai calcular area de um retangulo
func Rect(base, altura float64) float64 {
	return base * altura
}

func _TrianguloEq(base, altura float64) float64 {

	return base * altura / 2
}
