package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scanln(&n)

	sum := 0
	divisors := []int

	for i := 1; i < n; i++ {
		if n%i == 0 {
			divisors = append(divisors, i)
			sum += i
		}
	}

	divisorsStr := ""
	for i, divisor := range divisors {
		if i == len(divisors)-1 {
			divisorsStr += fmt.Sprintf("%d", divisor)
		} else {
			divisorsStr += fmt.Sprintf("%d + ", divisor)
		}
	}

	if sum == n {
		fmt.Printf("%d = %s = %d (NUMERO PERFEITO)\n", n, divisorsStr, n)
	} else {
		fmt.Printf("%d = %s = %d (NUMERO NAO E PERFEITO)\n", n, divisorsStr, sum)
	}
}
