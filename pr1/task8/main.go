package main

import (
	"fmt"
	"math"
)

func main() {
	var N int
	fmt.Scan(&N)

	for i := 0; math.Pow(2, float64(i)) <= float64(N); i++ {
		fmt.Println(math.Pow(2, float64(i)))
	}
}
