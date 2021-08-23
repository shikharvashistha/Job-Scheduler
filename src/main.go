package main 

import (
	"fmt"
	"time"

	"github.com/onatm/clockwerk"
)

type Job struct{}//{} is a struct literal

func (d Job) Run() {//this is the job
	time.Sleep(4 * time.Second)//sleep for 4 seconds
	fmt.Println("Every 30 seconds")//print every 30 seconds
}

func main() {
	var job Job//create a job
	c := clockwerk.New()//create a new clockwerk instance
	c.Every(30 * time.Second).Do(job)//every 30 seconds, do the job
	c.Start()//start the clockwerk instance
}
