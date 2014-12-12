package main

import (
	"fmt"
	"sync"
	"time"
)

type Job struct {
	i    int
	max  int
	text string
}

func outputText(job *Job, goGroup *sync.WaitGroup) {
	for job.i < job.max {
		time.Sleep(1 * time.Millisecond)
		fmt.Println(job.text)
		job.i++
	}
	goGroup.Done()
}

// Using a wait group to wait for all called goroutines to complete.
func main() {
	goGroup := new(sync.WaitGroup)
	fmt.Println("Starting")

	hello := new(Job)
	hello.text = "hello"
	hello.i = 0
	hello.max = 2

	world := new(Job)
	world.text = "world"
	world.i = 0
	world.max = 2

	go outputText(hello, goGroup)
	go outputText(world, goGroup)

	// 2 is the number of goroutines started.
	goGroup.Add(2)
	goGroup.Wait()
}
