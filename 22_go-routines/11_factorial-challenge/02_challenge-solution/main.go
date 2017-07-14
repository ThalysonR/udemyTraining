package main

import "fmt"

func main() {
	f := factorial(50)
	fmt.Println("Total:", <-f)
}

func factorial(n uint64) <-chan uint64 {
	return puller(decrementor(n))
}

func decrementor(n uint64) <-chan uint64 {
	out := make(chan uint64)
	go func() {
		for i := n; i > 0; i-- {
			out <- i
		}
		close(out)
	}()
	return out
}

func puller(c <-chan uint64) <-chan uint64 {
	out := make(chan uint64)
	go func() {
		sum := uint64(1)
		for n := range c {
			sum *= n
		}
		out <- sum
		close(out)
	}()
	return out
}
