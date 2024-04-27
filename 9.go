package main

import (
    "fmt"
    "math"
)

func main() {
    var N int

    fmt.Println("Digite um número inteiro positivo:")
    _, err := fmt.Scan(&N)

    // Verifica se ocorreu algum erro na leitura do número
    if err != nil {
        fmt.Println("Número inválido!")
        return
    }

    // Verifica se o número é positivo
    if N <= 0 {
        fmt.Println("Número inválido!")
        return
    }

    // Verifica se o número é primo
    if isPrime(N) {
        fmt.Println("PRIMO")
    } else {
        fmt.Println("NAO PRIMO")
    }
}

// Função para verificar se um número é primo
func isPrime(n int) bool {
    // Um número é primo se for maior que 1 e não tiver divisores além de 1 e ele mesmo
    if n <= 1 {
        return false
    }
    // Verifica divisores até a raiz quadrada de n
    for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
        if n%i == 0 {
            return false
        }
    }
    return true
}
