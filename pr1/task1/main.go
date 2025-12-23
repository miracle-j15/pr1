package main

import "fmt"

func main() {
	var x, sum int
	fmt.Print("Введите число: ")
	fmt.Scan(&x)

	for x > 0 {
		sum += x % 10
		x /= 10
	}

	fmt.Println("Сумма цифр", sum)
}
