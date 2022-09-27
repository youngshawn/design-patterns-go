package main

import "fmt"

type Fifo struct {
}

func (f *Fifo) evict(cache *Cache) {
	fmt.Println("Evicted by FIFO strategy")
}
