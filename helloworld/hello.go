package main

import (
	"fmt"
	//"github.com/itegel/gotest/stringutil"
)

const (
	A = 1
	B = 2
	C = iota
	D = iota
	E
)

func main() {
	fmt.Printf("Hello, world.\n")
	//fmt.Printf(stringutil.Reverse("I am itegel!\n"))
	fmt.Printf("C:%d, D:%d, E:%d\n", C, D, E)
	rating := map[int]int{100: 1, 200: 2, 300: 4}
	rate, ok := rating[200]
	fmt.Printf("200:%d, ok:%t", rate, ok)
}
