package main

import "fmt"

var (
	a int = 10
)

func foo(x float64) {
	fmt.Println(x)
}

func main() {
	foo(float64(a))
}
