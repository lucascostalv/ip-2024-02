package main

import "fmt"

func main() {
	// Entrada do usuário
	n := 4
	var resultado int

	if n == 0 {
		fmt.Println("0! = 1")
	}else{
		fat:=1
		for i:=1; i<=n; i++{
			fat = fat * i
		}
		resultado = fat
	}

	fmt.Printf("%d! = %d\n", n, resultado)
}
