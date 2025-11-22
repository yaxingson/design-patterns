package main

type Shape interface {
	paint()
}

type Circle struct{}

func (c *Circle) paint() {}

type Triangle struct{}

func (t *Triangle) paint() {}

type Rect struct{}

func (r *Rect) paint() {}

func paint(shape Shape) {
	shape.paint()
}

func init() {

}
