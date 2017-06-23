package main

import "fmt"

func zero(z *int) {
	fmt.Println(z)
	*z = 5
}

func main() {
	x := 5
	fmt.Println(&x)
	zero(&x)
	fmt.Println(x) // x is 0
}
