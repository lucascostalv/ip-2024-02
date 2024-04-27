package main

import (
	"fmt"
)

func main() {
	tamanho := 3
	for{
		if tamanho == 0{
			break
		}

		numeros := make([]float64, tamanho)
		for i := 0; i < tamanho; i++ {
			fmt.Scan(&numeros[i])
			for j:=0; j < tamanho; j++{
				if numeros[j] >= numeros[i+1]{
					fmt.Println("DESORDENADO")
				}else{
					fmt.Println("ORDENADO")
				}
			}
		}
	}
}