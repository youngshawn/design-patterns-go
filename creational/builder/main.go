package main

import "fmt"

func main() {
	nb := getBuilder("normal")
	ib := getBuilder("igloo")

	director := newDirector(nb)
	nHouse := director.buildHouse()
	fmt.Printf("%#v\n", nHouse)

	director.setBuilder(ib)
	iHouse := director.buildHouse()
	fmt.Printf("%#v\n", iHouse)
}
