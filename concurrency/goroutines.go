package main

import (
	"fmt"
	"time"
)

type Job struct {
	i    int
	max  int
	text string
}

func outputText(job *Job) {
	for job.i < job.max {
		time.Sleep(1 * time.Millisecond)
		fmt.Println(job.text)
		job.i++
	}
}

func main() {
	hello := new(Job)
	world := new(Job)

	hello.text = "hello"
	hello.i = 0
	hello.max = 3

	world.text = "world"
	world.i = 0
	world.max = 5

	go outputText(hello)
	outputText(world)
}
