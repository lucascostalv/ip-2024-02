package main

import "fmt"

func main() {
    var n, num, maxLength, currentLength int

    fmt.Scan(&n)

    // Lê o primeiro número
    fmt.Scan(&num)
    currentLength = 0
    maxLength = 0

    for i := 1; i < n; i++ {
        prevNum := num
        fmt.Scan(&num)

        // Verifica se o número atual é maior que o anterior
        if num > prevNum {
            currentLength++
        } else {
            // Se não for, verifica se o segmento atual é o maior até agora
            if currentLength > maxLength {
                maxLength = currentLength
            }
            // Reseta o comprimento atual do segmento crescente
            currentLength = 0
        }
    }

    // Verifica se o último segmento é o maior
    if currentLength > maxLength {
        maxLength = currentLength
    }

    // O comprimento do segmento é dado pelo número de elementos menos um
    fmt.Printf("O comprimento do segmento crescente maximo e: %d\n", maxLength)
}
