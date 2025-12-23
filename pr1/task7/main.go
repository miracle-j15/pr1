package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	switch {
	case n%10 == 1:
		fmt.Println(n, "korova")
	case n%10 > 4 || (n >= 11 && n <= 14) || n%10 == 0:
		fmt.Println(n, "korov")
	case n%10 > 1 && n%10 <= 4:
		fmt.Println(n, "korovy")

	}
}
