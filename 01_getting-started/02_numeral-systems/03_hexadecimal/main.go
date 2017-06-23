package main

import "fmt"

func main() {
	fmt.Printf("\n%d - %b - %x\n", 42, 42, 42)
	fmt.Printf("%d - %b - %#x\n", 42, 42, 42)
	fmt.Printf("%d - %b - %#X\n", 42, 42, 42)
}
