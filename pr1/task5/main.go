package main

import "fmt"

func sumc(n int) int {
	sum := 0

	for n > 0 {
		sum += n % 10
		n /= 10
	}

	return sum
}

func main() {
	var n int
	fmt.Scan(&n)

	for n > 10 {
		n = sumc(n)
	}

	fmt.Println(n)
}
