package main

import (
	"fmt"

	"github.com/lkeix/goleinu"
)

type User struct {
	ID int
	Name string
}

func main() {
	s, err := goleinu.New[[]*User](0, 10, goleinu.WithChunkSize(5), goleinu.WithMaxInMemorySize(10), goleinu.WithBufferSize(10))
	if err != nil {
		panic(err)
	}

	a, err := goleinu.Get[[]*User](s, 0)
	if err != nil {
		panic(err)
	}

	fmt.Println(a)
}