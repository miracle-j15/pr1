package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	num := b / 7 * 7
	if num >= a {
		fmt.Println(num)
	} else {
		fmt.Println("NO")
	}
}
