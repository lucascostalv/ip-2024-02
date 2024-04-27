package main

import (
    "fmt"
    "bufio"
    "os"
    "strings"
)

func main() {
    scanner := bufio.NewScanner(os.Stdin)

    for scanner.Scan() {
        line := scanner.Text()
        // Verifica se a linha contém apenas zeros, indicando o fim da entrada
        if line == "0 0 0.0" {
            break
        }

        // Divide a linha nos valores de matrícula, horas trabalhadas e valor por hora
        values := strings.Fields(line)
        matricula, horasTrabalhadas, valorPorHora := values[0], values[1], values[2]

        // Converte os valores de string para números inteiros ou flutuantes
        var matriculaInt int
        fmt.Sscanf(matricula, "%d", &matriculaInt)
        var horasTrabalhadasFloat, valorPorHoraFloat float64
        fmt.Sscanf(horasTrabalhadas, "%f", &horasTrabalhadasFloat)
        fmt.Sscanf(valorPorHora, "%f", &valorPorHoraFloat)

        // Calcula o salário do funcionário
        salario := horasTrabalhadasFloat * valorPorHoraFloat

        // Imprime a matrícula e o salário com duas casas decimais
        fmt.Printf("%d %.2f\n", matriculaInt, salario)
    }

    if err := scanner.Err(); err != nil {
        fmt.Fprintln(os.Stderr, "erro lendo a entrada:", err)
    }
}
