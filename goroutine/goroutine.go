package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	var a [10]int
	for i := 0; i < 10; i++ {
		//匿名函数
		go func(i int) {
			for {
				//fmt.Printf("goroutine print: %d\n", i)
				a[i]++
				runtime.Gosched()
			}
		}(i)
	}

	time.Sleep(time.Millisecond)
	fmt.Println(a)
}
