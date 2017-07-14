package main

import "fmt"

func wow() func(n int) (int, bool) {
	return func(n int) (int, bool) {
		dividido := n%2 == 0
		return n / 2, dividido
	}
}

func main() {
	half := wow()
	var n int
	fmt.Println("Informe um numero: ")
	fmt.Scan(&n)
	fmt.Println(half(n))
}
