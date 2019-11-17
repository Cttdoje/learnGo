package main

import (
	"bufio"
	"fmt"
	"learnGo/functional/fib"
	"os"
)

func tryDefer() {
	defer fmt.Println(1)
	fmt.Println(2)
}

func writeFile(fileName string) {
	file, err := os.OpenFile(fileName, os.O_EXCL|os.O_CREATE, 0666)
	if err != nil {
		if pathError, ok := err.(*os.PathError); !ok {
			fmt.Println("1111")
			panic(err)
		} else {
			fmt.Println("===")
			fmt.Println(pathError.Op,
				pathError.Path,
				pathError.Err)
		}
		fmt.Println("Error:", err.Error())
		return
	}
	//在程序结束前关闭，即使出现异常也会执行这句
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	f := fib.Fibonacci()
	for i := 0; i < 20; i++ {
		fmt.Fprintln(writer, f())
	}

}

func main() {
	writeFile("fib.txt")
}
