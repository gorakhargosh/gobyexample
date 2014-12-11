package main

import "fmt"

func foo(slice []string) {
	fmt.Println(slice)
}

func main() {
	a := [2]string{"hello", "world"}
	foo(a[:])
}
