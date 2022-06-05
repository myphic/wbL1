package main

import (
	"fmt"
	"math"
)

type Point struct {
	x, y float64
}

func (p *Point) NewPoint(x, y float64) {
	p.x = x
	p.y = y
}

func distance(a, b *Point) float64 {
	return math.Sqrt((b.x-a.x)*(b.x-a.x) + (b.y-a.y)*(b.y-a.y))
}

func main() {
	A := &Point{4, 2}
	B := &Point{1, 6}
	fmt.Println(distance(A, B))
}
