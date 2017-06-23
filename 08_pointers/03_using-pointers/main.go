package main

import "fmt"

func main() {
	a := 43

	fmt.Println(a)  // 43
	fmt.Println(&a) // 0x20818a220

	var b *int = &a
	fmt.Println(b) // 0x20818a220
	fmt.Println(*b)

	*b = 42        // b says, "The value at this address, change it to 42"
	fmt.Println(a) // 42
}
