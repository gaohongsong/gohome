package main

import (
	"fmt"
	"math"
)

type Circle struct {
	r    float64
	name string
}

func circleArea(c Circle) float64 {
	ret := math.Pi * c.r * c.r
	c.r = 100
	return ret
}

func circleArea1(c *Circle) float64 {
	ret := math.Pi * c.r * c.r
	c.r = 100
	return ret
}

// Methods
func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}

type Person struct {
	name string
}

func (p *Person) talk()  {
	fmt.Println("Hi, my name is", p.name)
}

// Embedded types
type Boss struct {
	Person
	company string
}

func (b *Boss) talk() {
	fmt.Println("Hi, i am boss of", b.company)
}

func main() {
	var c Circle
	//{0}
	fmt.Println(c)
	c.r = 123
	c.name = "c"
	//{123 c}
	fmt.Println(c)

	//&{0 c1} a pointer to circle
	c1 := new(Circle)
	//c11 := &Circle{0, "c1"}
	c1.name = "c1"
	fmt.Println(c1)

	//{911} {912}
	c2 := Circle{r: 911, name: "c2"}
	c3 := Circle{912, "c3"}
	fmt.Println(c2, c3)
	//circleArea: 47529.15525615998 -> r = 123
	fmt.Printf("circleArea: %v -> r = %v\n", circleArea(c), c.r)
	//circleArea1: 47529.15525615998 -> r = 100 changed
	fmt.Printf("circleArea1: %v -> r = %v\n", circleArea1(&c), c.r)
	//c: 31415.926535897932 or 31415.926535897932, c2: 2.6072737166598947e+06, c3: 2.613000840067389e+06
	fmt.Printf("c: %v or %v, c2: %v, c3: %v\n", c.area(), (&c).area(), c2.area(), c3.area())

	b := Boss{Person{"miya"}, "tencent"}
	//Hi, i am boss of tencent
	b.talk()
	//Hi, my name is miya
	b.Person.talk()
	fmt.Printf("b.name: %v\n", b.name)
}
