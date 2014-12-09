package main

import "fmt"

const (
	StatusOK                   = 200
	StatusCreated              = 201
	StatusAccepted             = 202
	StatusNonAuthoritativeInfo = 203
	StatusNoContent            = 204
	StatusResetContent         = 205
	StatusPartialContent       = 206
)

const (
	Pi    = 3.14
	Truth = false
	Big   = 1 << 100
	Small = Big >> 99
)

func main() {
	const Greeting = "asldfasdf"
	fmt.Println(Greeting)
	fmt.Println(Pi)
	fmt.Println(Truth)
}
