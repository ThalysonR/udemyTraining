package main

import "fmt"

func main() {
	var nome string
	fmt.Print("Digite seu nome: ")
	fmt.Scan(&nome)
	fmt.Println("Olá,", nome)
}
