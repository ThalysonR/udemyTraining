package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func menor(n1 string, n2 string) int {
	if len(n1) < len(n2) {
		return len(n1)
	} else {
		return len(n2)
	}
}

func main() {
	var names []string
	var m int
	names = make([]string, 0)
	file, err := ioutil.ReadFile("./names.txt")
	check(err)

	for _, n := range strings.Split(string(file), "\"") {
		if n != "," {
			names = append(names, n)
		}
	}

	for i := 0; i < len(names)-1; i++ {
		for j := 1; j < len(names); j++ {
			m = menor(names[i], names[j])
			for k := 0; k < m; k++ {
				if byte(names[i][k]) < byte(names[j][k]) {
					names[i], names[j] = names[j], names[i]
					break
				} else if byte(names[i][k]) > byte(names[j][k]) {
					break
				} else if k == m-1 {
					if len(names[j]) > len(names[i]) {
						names[i], names[j] = names[j], names[i]
					}
				}
			}
		}
	}
	var sum int
	var total int64
	for i, n := range names {
		sum = 0
		for _, lt := range n {
			fmt.Println(string(lt), " ", lt-64)
			sum += int(lt) - 64
		}
		fmt.Println(i, " * ", sum, " = ", i*sum)
		total += int64(sum * i)
		fmt.Println("Total: ", total)
	}
}
