package goarea

import "math"

// Pi é um número ... (Tem que fazer pq é público)
const Pi = 3.141596

// Circ é responsável por calcular a área da circunferencia
func Circ(raio float64) float64 {
	return Pi * math.Pow(raio, 2)
}

// Rect é responsável por calcular a área de um retangulo
func Rect(base, altura float64) float64 {
	return base * altura
}

// Não é uma função visivel por causa do "_" na frente
func _TrianguloEq(lado float64) float64 {
	return ((math.Pow(lado, 2) * math.Sqrt(3)) / 4)
}
