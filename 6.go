package main

import "fmt"

func main() {
    var n, num, maxCount, currentCount int

    fmt.Scan(&n)

    // Lê o primeiro número
    fmt.Scan(&num)
    currentCount = 1
    maxCount = 1

    for i := 1; i < n; i++ {
        prevNum := num
        fmt.Scan(&num)

        // Verifica se o número atual é igual ao anterior
        if num == prevNum {
            currentCount++
        } else {
            // Se não for, verifica se a subsequência atual é a maior até agora
            if currentCount > maxCount {
                maxCount = currentCount
            }
            // Reseta o contador da subsequência atual
            currentCount = 1
        }
    }

    // Verifica se a última subsequência é a maior
    if currentCount > maxCount {
        maxCount = currentCount
    }

    // Imprime o resultado
    fmt.Printf("A maior subsequência de elementos iguais possui %d elementos.\n", maxCount)
}
