package main

import (
	"fmt"
)

func main() {
	a := 43

	fmt.Println(a)
	fmt.Println(&a)

	var b *int = &a

	fmt.Println(&b)
}

// the above code makes b a ponter to the memory address where a is stored
// b is of type "int pointer"
// *int -- the * is part of the type -- b is of type *int
