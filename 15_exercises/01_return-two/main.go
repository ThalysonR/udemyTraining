package main

import "fmt"

func half(n int) (int, bool) {
	dividido := n%2 == 0
	return n / 2, dividido
}

func main() {
	var n int
	fmt.Println("Informe um numero: ")
	fmt.Scan(&n)
	fmt.Println(half(n))
}
