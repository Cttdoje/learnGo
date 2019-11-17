package main

import (
	"fmt"
)

func main() {
	chanDemo()
}

type worker struct {
	in   chan int
	done chan bool
}

// 读取channel中数据
func doWorker(id int, c chan int, done chan bool) {
	for n := range c {
		fmt.Printf("worker %d received %d\n", id, n)
		done <- true
	}

}

func createWorker(id int) worker {
	w := worker{
		in:   make(chan int),
		done: make(chan bool),
	}
	go doWorker(id, w.in, w.done)
	return w
}

func chanDemo() {
	var workers [10]worker
	for i := 0; i < 10; i++ {
		workers[i] = createWorker(i)
	}

	for i := 0; i < 10; i++ {
		workers[i].in <- 'a' + i
		<-workers[i].done
	}
}
