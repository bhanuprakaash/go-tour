package main

import (
	"fmt"
	"math"
)

type Point struct{ X, Y float64 }
type Path []Point
type P struct{ X, Y float64 }

func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

func (path *Path) Distance() float64 {
	sum := 0.0
	points := *path
	for i := range points {
		if i > 0 {
			sum += points[i-1].Distance(points[i])
		}
	}
	return sum
}

func (p *P) Scaleby(factor float64) {
	p.X *= factor
	p.Y *= factor
}

func main() {
	p := Point{3, 2}
	q := Point{2, 3}
	fmt.Println(p.Distance(q))
	fmt.Println(Point{2, 3}.Distance(q))

	perim := &Path{
		{1, 1},
		{5, 1},
		{5, 4},
		{1, 1},
	}
	fmt.Println(perim.Distance())

	pointP := &P{1, 2}
	pointP.Scaleby(2)
	fmt.Println(*pointP)

}
