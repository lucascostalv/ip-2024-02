package main

import (
	"fmt"
)

func main() {
	var m,n int
	fmt.Scanln(&m,&n)
	for i := 0; i < m; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d,%d ",i,j)
			if j < i -1 {
				fmt.Print("-")
			}
		}
		fmt.Println()
	}	
}