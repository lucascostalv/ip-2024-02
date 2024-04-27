package main

import (
	"fmt"
)

func main() {
	var (
		qtTime, int 
		// time1,time2,time3 int
	)	
	qtTime = 3
	// calcular 
	fmt.Println("Digite a quantidade de times")
	// fmt.Scan(&n)
	if qtTime <= 2{
		fmt.Println("Quantidade de times invalida")
	}else{
		k := 1
		for i := 1; i < qtTime; i++ {
			for j := i + 1; j < qtTime; j++ {
				//para gerar todaas as combinacoes possiveis, o time 1 nao pode ser igual ao time 2~
				//time 1 nao pode ser igual ao time 3 e o time 2 nao pode ser igual ao time 3
				fmt.Printf("Final %d time %d x time %d\n",k, i, j)
			}
		}
	}
}