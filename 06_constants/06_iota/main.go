package main

import "fmt"

const (
	_  = iota             // 0
	KB = 1 << (iota * 10) // 1 << (1 * 10)
	MB = 1 << (iota * 10) // 1 << (2 * 10)
	GB = 1 << (iota * 10) // 1 << (3 * 10)
)

func main() {
	a := 65535
	a = a >> 5 << 5
	// a = a << 24
	b := a >> 1
	fmt.Println("Binary\t\tDecimal")
	fmt.Printf("%b\t", KB)
	fmt.Printf("%d\n", KB)
	fmt.Printf("%b\t", MB)
	fmt.Printf("%d\n", MB)
	fmt.Printf("%b\t", GB)
	fmt.Printf("%d\n", GB)
	fmt.Printf("%b\t", a)
	fmt.Printf("%d\n", a)
	fmt.Printf("%b\t", b)
	fmt.Printf("%d\n", b)
}
