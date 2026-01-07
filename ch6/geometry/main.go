package main

import (
	"fmt"
	"math"
	"net/url"
)

type Point struct{ X, Y float64 }

func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

type Path []Point

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

type P struct{ X, Y float64 }

func (p *P) Scaleby(factor float64) {
	p.X *= factor
	p.Y *= factor
}

type Intlist struct {
	Value int
	Tail  *Intlist
}

func (list *Intlist) Sum() int {
	if list == nil {
		return 0
	}

	return list.Value + list.Tail.Sum()
}

type Values map[string][]string

func (v Values) Get(key string) string {
	if vs := v[key]; len(vs) > 0 {
		return vs[0]
	}
	return ""
}

func (v Values) Add(key, value string) {
	v[key] = append(v[key], value)
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

	item1 := Intlist{
		Value: 34,
		Tail:  nil,
	}
	item2 := Intlist{
		Value: 23,
		Tail:  &item1,
	}
	fmt.Println(item2.Sum())

	m := url.Values{"lang": {"en"}}
	m.Add("item", "1")
	m.Add("item", "2")
	fmt.Println(m.Get("lang"))
	fmt.Println(m.Get("q"))
	fmt.Println(m.Get("item"))
	fmt.Println(m["item"])
	// m = nil
	// fmt.Println(m.Get("item"))
	// m.Add("item", "3")

}
