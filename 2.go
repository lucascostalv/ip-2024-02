package main

import "fmt"

func main() {
    // Passo 1: Ler as populações dos países A e B da entrada
    var populacaoA, populacaoB int
    fmt.Scanln(&populacaoA)
    fmt.Scanln(&populacaoB)

    // Passo 2: Definir as taxas de crescimento para cada país
    taxaCrescimentoA := 0.03
    taxaCrescimentoB := 0.015

    // Passo 3: Calcular o número de anos necessários
    anos := 0
    for populacaoA <= populacaoB {
        populacaoA = int(float64(populacaoA) * (1 + taxaCrescimentoA))
        populacaoB = int(float64(populacaoB) * (1 + taxaCrescimentoB))
        anos++
    }

    // Passo 4: Exibir o resultado na saída
    fmt.Printf("ANOS = %d\n", anos)
}
