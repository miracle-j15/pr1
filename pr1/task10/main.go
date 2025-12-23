package main

import "fmt"

func main() {
	var N, d, res string
	fmt.Scan(&N, &d)

	for i := 0; i < len(N); i++ {
		if string(N[i]) != d {
			res += string(N[i])
		}
	}

	fmt.Println(res)
}
