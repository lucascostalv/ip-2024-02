package main

import (
	"fmt"
)

func main() {
	// Inicialização das variáveis
	var (
		sumPar, sumImpar float64
		countPar, countImpar int
		num float64
	)

	// Entrada de dados
	fmt.Println("Insira uma sequência de números inteiros diferentes de zero (digite 0 para finalizar):")
	for {
		fmt.Scan(&num)
		if num == 0 {
			break // Encerra o loop quando o usuário digita 0
		}
		if int(num)%2 == 0 {
			sumPar += num
			countPar++
		} else {
			sumImpar += num
			countImpar++
		}
	}

	// Cálculo das médias
	mediaPar := sumPar / float64(countPar)
	mediaImpar := sumImpar / float64(countImpar)

	// Saída dos resultados
	fmt.Printf("MEDIA PAR: %.2f\n", mediaPar)
	fmt.Printf("MEDIA IMPAR: %.2f\n", mediaImpar)
}
