package main

import "fmt"

func main() {
	name := "Sydney"
	str, ok := name.(string)
	if ok {
		fmt.Print("%q\n", str)
	} else {
		fmt.Printf("value is not a string\n")
	}
}
