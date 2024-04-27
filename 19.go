package main

import (
	"fmt"
)

func main() {
	var m int
	fmt.Print("Digite um número inteiro maior que zero: ")
	fmt.Scan(&m)

	for n := 1; n <= m; n++ {
		target := n * n * n
		var numero_impar int
		sum := 0
		num := 1
		for sum < target {
			numero_impar += num
			sum += num
			num += 2
		}
		formatted := ""
		for i := 0; i < n; i++ {
			if i > 0 {
				formatted += " + "
			}
			formatted += fmt.Sprintf("%d", numero_impar)
			numero_impar += 2
		}
		fmt.Printf("%d * %d * %d = %s\n", n, n, n, formatted)
	}
}
