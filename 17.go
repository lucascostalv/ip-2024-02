package main

import (
	"fmt"
	"math"
)

type Mercadoria struct {
	Codigo       uint64
	PrecoCompra  float64
	PrecoVenda   float64
	NumVendas    int
}

func main() {
	var mercadorias []Mercadoria

	// Ler as mercadorias
	for {
		var codigo uint64
		var precoCompra, precoVenda float64
		var numVendas int

		_, err := fmt.Scan(&codigo, &precoCompra, &precoVenda, &numVendas)
		if err != nil {
			break
		}

		mercadoria := Mercadoria{Codigo: codigo, PrecoCompra: precoCompra, PrecoVenda: precoVenda, NumVendas: numVendas}
		mercadorias = append(mercadorias, mercadoria)
	}

	// Variáveis para contagem e cálculos
	lucroMenor10 := 0
	lucroEntre10e20 := 0
	lucroMaior20 := 0
	maiorLucro := Mercadoria{}
	maisVendida := Mercadoria{}
	totalCompra := 0.0
	totalVenda := 0.0

	// Calcular e analisar as mercadorias
	for _, mercadoria := range mercadorias {
		lucro := ((mercadoria.PrecoVenda - mercadoria.PrecoCompra) / mercadoria.PrecoCompra) * 100
		totalCompra += mercadoria.PrecoCompra * float64(mercadoria.NumVendas)
		totalVenda += mercadoria.PrecoVenda * float64(mercadoria.NumVendas)

		if lucro < 10 {
			lucroMenor10++
		} else if lucro <= 20 {
			lucroEntre10e20++
		} else {
			lucroMaior20++
		}

		if lucro > ((maiorLucro.PrecoVenda - maiorLucro.PrecoCompra) / maiorLucro.PrecoCompra * 100) {
			maiorLucro = mercadoria
		}

		if mercadoria.NumVendas > maisVendida.NumVendas {
			maisVendida = mercadoria
		}
	}

	// Calcular o percentual de lucro total
	lucroTotal := ((totalVenda - totalCompra) / totalCompra) * 100

	// Imprimir resultados
	fmt.Printf("Quantidade de mercadorias que geraram lucro menor que 10%%: %d\n", lucroMenor10)
	fmt.Printf("Quantidade de mercadorias que geraram lucro maior ou igual a 10%% e menor ou igual a 20%%: %d\n", lucroEntre10e20)
	fmt.Printf("Quantidade de mercadorias que geraram lucro maior do que 20%%: %d\n", lucroMaior20)
	fmt.Printf("Codigo da mercadoria que gerou maior lucro: %d\n", maiorLucro.Codigo)
	fmt.Printf("Codigo da mercadoria mais vendida: %d\n", maisVendida.Codigo)
	fmt.Printf("Valor total de compras: %.2f, valor total de vendas: %.2f e percentual de lucro total: %.2f%%\n", totalCompra, totalVenda, lucroTotal)
}
