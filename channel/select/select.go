package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generator() chan int {
	out := make(chan int)
	go func() {
		i := 0
		for {
			time.Sleep(time.Duration(rand.Intn(1500)) * time.Millisecond)
			out <- i
			i++
		}
	}()
	return out
}

func worker(id int, c chan int) {
	for n := range c {
		time.Sleep(time.Second)
		fmt.Printf("worker %d received %d\n", id, n)
	}
}

func createWorker(id int) chan<- int {
	c := make(chan int)
	go worker(id, c)
	return c
}

func main() {
	var c1, c2 = generator(), generator()
	var worker = createWorker(0)

	//var w chan int  // w == nil
	//n := 0
	//hasValue := false
	var values []int

	chanTimes := time.After(10 * time.Second)

	tick := time.Tick(time.Second)

	for {
		// chan<- 只写   <-chan 只读
		var activeWorker chan<- int
		var activeValue int
		if len(values) > 0 {
			activeWorker = worker
			activeValue = values[0]
		}

		select {
		case n := <-c1:
			values = append(values, n)
		case n := <-c2:
			values = append(values, n)
		case activeWorker <- activeValue: //  activeWorker如果是nil,则不会被case到
			values = values[1:]
		case <-tick:
			fmt.Println("values len = ", len(values))
		case <-chanTimes:
			fmt.Println("bye bye")
			return
		}
	}
}
