package main

import (
	"fmt"
	"learngo/interface/retriever/mock"
	"learngo/interface/retriever/real"
	"time"
)

const url = "http://www.oktfolio.me"

func main() {

	var r Retriever
	retriever := &mock.Retriever{"this is a fake oktfolio.me"}

	r = retriever
	inspect(r)

	r = &real.Retriever{
		UserAgent: "Mozilla/5.0",
		TimeOut:   time.Minute,
	}
	inspect(r)

	// Type assertion
	realRetriever := r.(*real.Retriever)
	fmt.Println(realRetriever)
	mockRetriever, ok := r.(*mock.Retriever)
	if ok {
		fmt.Println(mockRetriever)
	} else {
		fmt.Println("not a mock retriever")
	}

	fmt.Println("Try a session")
	fmt.Println(session(retriever))
}

func inspect(r Retriever) {
	fmt.Println("Inspecting", r)
	fmt.Printf(" > %T %v\n", r, r)
	fmt.Print(" > Type switch:")
	switch v := r.(type) {
	case *mock.Retriever:
		fmt.Println("Contents : ", v.Contents)
	case *real.Retriever:
		fmt.Println("UserAgent : ", v.UserAgent)
	}
	fmt.Println()
}

type Poster interface {
	Post(url string, form map[string]string) string
}

func download(r Retriever) string {
	return r.Get(url)
}

func post(poster Poster) {
	poster.Post(url, map[string]string{
		"name":   "oktfolio",
		"course": "golang",
	})
}

type RetrieverPoster interface {
	Retriever
	Poster
}

func session(poster RetrieverPoster) string {
	poster.Post(url, map[string]string{
		"contents": "another faked oktfolio.me",
	})
	return poster.Get(url)
}

type Retriever interface {
	Get(url string) string
}
