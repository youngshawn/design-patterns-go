package main

type Rectangle struct {
	l int
	b int
}

func (t *Rectangle) accpet(v Visitor) {
	v.visitForRectangle(t)
}

func (t *Rectangle) getType() string {
	return "Rectangle"
}
