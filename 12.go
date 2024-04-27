package main

import (
	"fmt"
)

func main() {
	var valorIngresso, valorInicial, valorFinal float64

	fmt.Println("Digite o valor do ingresso:")
	fmt.Scanln(&valorIngresso)

	fmt.Println("Digite o valor inicial do intervalo:")
	fmt.Scanln(&valorInicial)

	fmt.Println("Digite o valor final do intervalo:")
	fmt.Scanln(&valorFinal)

	if valorInicial >= valorFinal {
		fmt.Println("INTERVALO INVALIDO.")
		return
	}

	maxLucro := 0.0
	qtdIngressosMax := 0
	melhorValor := 0.0

	for valor := valorInicial; valor <= valorFinal; valor += 1.0 {
		qtdIngressos := calcularIngressos(valor, valorIngresso)
		lucro := calcularLucro(valor, qtdIngressos, valorIngresso)

		fmt.Printf("V: %.2f, N: %d, L: %.2f\n", valor, qtdIngressos, lucro)

		if lucro > maxLucro {
			maxLucro = lucro
			qtdIngressosMax = qtdIngressos
			melhorValor = valor
		}
	}

	fmt.Printf("\nMelhor valor final: %.2f\n", melhorValor)
	fmt.Printf("Lucro: %.2f\n", maxLucro)
	fmt.Printf("Numero de ingressos: %d\n", qtdIngressosMax)
}

func calcularIngressos(valor, valorIngresso float64) int {
	if valor < valorIngresso {
		return 120
	} else if valor > valorIngresso {
		return 120 + int((valor - valorIngresso) / 0.5) * 25
	}
	return 120
}

func calcularLucro(valor float64, qtdIngressos int, valorIngresso float64) float64 {
	despesasFixas := 200.0 + (0.05 * float64(qtdIngressos))
	receita := valor * float64(qtdIngressos)
	return receita - despesasFixas
}
