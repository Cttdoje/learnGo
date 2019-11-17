package main

import (
	"errors"
	"fmt"
)

func main() {
	tryRecover()
}

func tryRecover() {
	defer func() {
		r := recover()
		if err, ok := r.(error); ok {
			fmt.Println("error occurred :", err)
		} else {
			panic(err)
		}
	}()
	panic(errors.New("this my panic"))
}
