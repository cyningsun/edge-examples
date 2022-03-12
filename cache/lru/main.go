package main

import (
	"fmt"

	"github.com/cyningsun/edge"
	"github.com/cyningsun/edge/cache"
)

type Data struct {
	val string
}

func main() {
	var c edge.Cache
	c, err := cache.NewLRU()
	if err != nil {
		fmt.Printf("NewLRU error:%v\n", err)
		return
	}
	want := Data{"val1"}
	c.Set("key1", want)
	got, ok := c.Get("key1")
	if !ok {
		return
	}
	if got != want {
		fmt.Printf("LRU error got:%v, want:%v\n", got, want)
	}
	fmt.Printf("LRU got:%v, want:%v\n", got, want)
}
