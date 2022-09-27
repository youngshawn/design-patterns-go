package main

import "fmt"

type Lfu struct {
}

func (l *Lfu) evict(cache *Cache) {
	fmt.Println("Evicted by LFU strategy")
}
