package main

import (
	"fmt"
	"learnGo/retriever/mock"
	"learnGo/retriever/real"
	"time"
)

type Retriever interface {
	Get(url string) string
}

type Poster interface {
	Post(url string, form map[string]string) string
}

//接口组合
type RetrieverPoster interface {
	Retriever
	Poster
}

func download(r Retriever) string {
	return r.Get("http://www.91mika.com")
}

func post(p Poster) {
	p.Post("http://www.91mika.com", map[string]string{
		"name":   "cttdoje",
		"course": "golang",
	})
}

//判断实现类型
func inspect(r Retriever) {
	switch valueType := r.(type) {
	case mock.Retriever:
		fmt.Println("contents:", valueType.Contents)
	case *real.Retriever:
		fmt.Println("userAgent", valueType.UserAgent)
	}
}

func main() {
	var r Retriever
	r = mock.Retriever{"hello cttdoje"}
	r = &real.Retriever{
		UserAgent: "Mozilla/5.0",
		TimeOut:   time.Minute,
	}
	fmt.Printf("%T %v\n", r, r)
	//fmt.Println(download(r))

	inspect(r)

	// Type assertion
	realRetriever := r.(*real.Retriever)
	fmt.Println(realRetriever.TimeOut)

}
