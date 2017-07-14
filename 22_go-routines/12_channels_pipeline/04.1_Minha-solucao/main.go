package main

import (
	"fmt"
	"time"
)

var done = make(chan bool)

func main() {
	c := make(chan uint64)

	go func() {
		for i := 0; i < 100; i++ {
			factorial(uint64(i), c)
		}
	}()

	go func() {
		<- done
		close(c)
	}()

	for n := range c {
		fmt.Println(n)
	}
}

func factorial(fac uint64, c chan uint64) {
	go func() {
		total := uint64(1)
		for i := fac; i > 0; i-- {
			total *= i
		}
		c <- total
		if fac == 99 {
			time.Sleep(2 * time.Millisecond)
			done <- true
		}
	}()
}
