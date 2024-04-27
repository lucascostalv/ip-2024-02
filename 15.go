package main

import (
	"fmt"
)

func main() {
	var t int
	var reversaoA, reversaoB int
	var reversao int
	fmt.Scan(&t)
	for i := 0; i < t; i++ {
		var a,b int	
		fmt.Scan(&a,&b)
		if t > 0 {
			digito:= a % 10	
			reversao = reversao*10 + digito	
			a = a / 10	
		}
		if reversaoA > reversaoB {
			fmt.Println("%d",reversaoA)
		} else {
			fmt.Println("%d",reversaoB)
		}
	}
}