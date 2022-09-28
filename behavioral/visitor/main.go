package main

import "fmt"

func main() {
	square := &Square{side: 2}
	circle := &Circle{radius: 3}
	rectangle := &Rectangle{l: 2, b: 3}

	areaCalculator := &AreaCalculator{}

	square.accept(areaCalculator)
	circle.accept(areaCalculator)
	rectangle.accpet(areaCalculator)

	fmt.Println()

	middleCoordinates := &MiddleCoordinates{}
	square.accept(middleCoordinates)
	circle.accept(middleCoordinates)
	rectangle.accpet(middleCoordinates)

	//middleCoordinates.visitForSquare(square)
	//middleCoordinates.visitForCircle(circle)
	//middleCoordinates.visitForRectangle(rectangle)
}
