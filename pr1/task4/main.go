package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	var num int
	fmt.Scan(&num)
	count := 1
	min := num

	for i := 1; i < n; i++ {
		fmt.Scan(&num)
		if num < min {
			min = num
			count = 1
		} else if num == min {
			count += 1
		}
	}

	fmt.Println(count)
}
