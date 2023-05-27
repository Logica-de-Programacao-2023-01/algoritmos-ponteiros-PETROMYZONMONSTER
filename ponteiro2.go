package main

import "fmt"

func main() {
	n1 := 5
	var ptr *int = &n1
	atualizar(ptr)
	fmt.Println(n1)
}

func atualizar(ptr *int) int {
	if *ptr%2 != 0 {
		*ptr = 1
	} else {
		*ptr = 0
	}
	return *ptr
}
