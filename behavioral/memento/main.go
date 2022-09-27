package main

import "fmt"

func main() {
	caretaker := &CareTaker{
		mementoArray: make([]*Memento, 0),
	}

	originator := &Originator{
		state: "A",
	}

	fmt.Printf("Originator Current state: %s\n", originator.getState())
	caretaker.addMemento(originator.createMemento())

	originator.seteState("B")
	fmt.Printf("Originator Current state: %s\n", originator.getState())
	caretaker.addMemento(originator.createMemento())

	originator.seteState("C")
	fmt.Printf("Originator Current state: %s\n", originator.getState())
	caretaker.addMemento(originator.createMemento())

	originator.restoreMemento(caretaker.getMemento(1))
	fmt.Printf("Restore to State: %s\n", originator.getState())

	originator.restoreMemento(caretaker.getMemento(0))
	fmt.Printf("Restore to State: %s\n", originator.getState())
}
