package main

import "fmt"

func greatest(arr ...int) int {
	var gt int
	for _, n := range arr {
		if n > gt {
			gt = n
		}
	}
	return gt
}

func main() {
	var slc []int
	//slc = make([]int,0)
	for i := 0; ; {
		fmt.Println("Informe um número (-1 para sair) : ")
		fmt.Scan(&i)
		if i != -1 {
			slc = append(slc, i)
		} else {
			break
		}
	}
	fmt.Println(greatest(slc...))
}
