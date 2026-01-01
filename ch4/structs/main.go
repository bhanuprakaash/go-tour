package main

import (
	"fmt"
	"time"
)

type Employee struct {
	ID            int
	Name, Address string
	DoB           time.Time
	Position      string
	Salary        int
	ManagerId     int
}

type address struct {
	hostname string
	port     int
}

type Wheel struct {
	Circle
	Spokes int
}

func main() {

	var dilbert = Employee{
		ID:        1,
		Name:      "Dilber",
		Address:   "Bangalore",
		DoB:       time.Date(1990, 5, 10, 0, 0, 0, 0, time.UTC),
		Position:  "Engineer",
		Salary:    54000,
		ManagerId: 2,
	}
	var employeeOfMonth *Employee = &dilbert

	// employeeOfMonth.Position += "(Pro active Player)"
	(*employeeOfMonth).Position += "(Pro active Player)"

	dilbert.Salary += 500

	fmt.Println(dilbert)
	fmt.Println(*employeeOfMonth)
	Bonus(&dilbert, 21)

	hits := make(map[address]int)
	hits[address{"golang.org", 443}]++

	fmt.Println(hits[address{"golang.org", 443}])

	var w Wheel

	// w.X = 1
	// w.Y = 2
	// w.Spokes = 23
	// w.Radius = 23

	w = Wheel{
		Circle: Circle{
			Point:  Point{X: 8, Y: 8},
			Radius: 5,
		},
		Spokes: 20,
	}

	fmt.Printf("%#v\n", w)

}

func Bonus(e *Employee, percent int) int {
	return e.Salary * percent / 100
}
