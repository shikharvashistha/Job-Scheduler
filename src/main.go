package main

import (
	"fmt"
	"time"

	"github.com/onatm/clockwerk"
)

type Job struct{}

func (d Job) Run() {
	time.Sleep(4 * time.Second)
	fmt.Println("Every 30 seconds")
}

func main() {
	var job Job
	c := clockwerk.New()
	c.Every(30 * time.Second).Do(job)
	c.Start()
}
