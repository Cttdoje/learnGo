package main

import (
	"fmt"
	"golang.org/x/tools/container/intsets"
	"learnGo/queue"
)

func testSparse() {
	s := intsets.Sparse{}

	s.Insert(1)
	s.Insert(1000)
	s.Insert(100000)
	fmt.Println(s.Has(1000))
	fmt.Println(s.Has(10000))
}

func main() {
	q := queue.Queue{1}
	q.Push(3)
	q.Push(2)
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.IsEmpty())
	fmt.Println(q.Pop())
	fmt.Println(q.IsEmpty())

	q.Push("abc")
	fmt.Println(q.IsEmpty())
	fmt.Println(q.Pop())

	/*fmt.Println("")
	testSparse()*/
}
