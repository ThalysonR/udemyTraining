package main

import "fmt"

func main() {
	var soma int
	for i := 1; i < 1000; i++ {
		if i%3 == 0 || i%5 == 0 {
			soma += i
		}
	}
	fmt.Println("Soma:", soma)
}
