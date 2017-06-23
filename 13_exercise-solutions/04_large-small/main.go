package main

import "fmt"

func main() {
	var a, b int32
	fmt.Print("Digite dois numeros, um pequeno e um maior: ")
	fmt.Scan(&a, &b)
	fmt.Println("Resto:", b%a)
}
