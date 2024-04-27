package main

import (
	"fmt"
)

func main() {
	var n int 
	fmt.Scan(&n)
	totalGraos := n
	for i := 1; i < n; i++ {
		if i % 2 == 0 {
			totalGraos += n * 2
		} else{
			totalGraos += n
		}
		n = n * 2
		}	
	fmt.Println(totalGraos)
}