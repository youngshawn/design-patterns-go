package main

import "fmt"

type Lru struct {
}

func (l *Lru) evict(cache *Cache) {
	fmt.Println("Evicted by LRU strategy")
}
