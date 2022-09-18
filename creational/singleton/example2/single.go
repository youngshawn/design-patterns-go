package main

import (
	"fmt"
	"sync"
	"time"
)

var once sync.Once

type single struct {
}

var singleInstance *single

func getInstance() *single {
	if singleInstance == nil {
		once.Do(
			func() {
				fmt.Println("Creating single instance now.")
				time.Sleep(time.Second)
				singleInstance = &single{}
			})
		fmt.Println("2. Single instance already created.")
	} else {
		fmt.Println("1. Single instance already created.")
	}

	return singleInstance
}
