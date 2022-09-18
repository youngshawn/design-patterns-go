package main

import (
	"fmt"
	"sync"
	"time"
)

var lock = &sync.Mutex{}

type single struct{}

var singleInstance *single

func getInstance() *single {
	if singleInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if singleInstance == nil {
			fmt.Println("Create single instance now.")
			time.Sleep(time.Second)
			singleInstance = &single{}
		} else {
			fmt.Println("2. Single instance already created.")
		}
	} else {
		fmt.Println("1. Single instance already created.")
	}
	return singleInstance
}
