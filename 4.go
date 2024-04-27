package main

import (
	"fmt"
)

func main() {
	// Leitura dos números fornecidos pelo usuário
	var n, i, K int
	var s float64

	fmt.Println("Digite um número de 0 a 9:")
	fmt.Scanln(&n)

	fmt.Println("Digite o valor inicial i:")
	fmt.Scanln(&i)

	fmt.Println("Digite a quantidade de valores K:")
	fmt.Scanln(&K)

	fmt.Println("Digite o incremento s:")
	fmt.Scanln(&s)

	// Tabuada de soma
	fmt.Println("Tabuada de soma:")
	for j := 0; j < K; j++ {
		B := float64(i) + float64(j)*s
		resultado := float64(n) + B
		fmt.Printf("%.2f + %.2f = %.2f\n", float64(n), B, resultado)
	}

	// Tabuada de subtração
	fmt.Println("Tabuada de subtração:")
	for j := 0; j < K; j++ {
		B := float64(i) + float64(j)*s
		resultado := float64(n) - B
		fmt.Printf("%.2f - %.2f = %.2f\n", float64(n), B, resultado)
	}

	// Tabuada de multiplicação
	fmt.Println("Tabuada de multiplicação:")
	for j := 0; j < K; j++ {
		B := float64(i) + float64(j)*s
		resultado := float64(n) * B
		fmt.Printf("%.2f x %.2f = %.2f\n", float64(n), B, resultado)
	}

	// Tabuada de divisão
	fmt.Println("Tabuada de divisão:")
	for j := 0; j < K; j++ {
		B := float64(i) + float64(j)*s
		resultado := float64(n) / B
		fmt.Printf("%.2f / %.2f = %.2f\n", float64(n), B, resultado)
	}
}
