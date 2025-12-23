package main

import "fmt"

func main() {
	var A int
	fmt.Scan(&A)
	f1, f2 := 1, 1
	n := 2
	for f2 < A {
		f1, f2 = f2, f1+f2
		n++
	}

	if f2 == A {
		fmt.Println(n)
	} else {
		fmt.Println(-1)
	}
}
