package main

import (
	"fmt"
	"math"
)

func main() {
	var n,e float64
	fmt.Scan(&n)
	fmt.Scan(&e)
	r :=1.0	
	for{
		nextR := 0.5*(r + n/r)
		erro := math.Abs(n-nextR*nextR)
		fmt.Println("r: %.9f, erro: %.9f\n",nextR,erro)
		if erro < e{
			break
		}
		r = nextR
	}

}