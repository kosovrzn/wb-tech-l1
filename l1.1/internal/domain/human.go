package domain

import "fmt"

type Human struct {
	Name string
	Age  int
}

func (h Human) Greet() {
	fmt.Printf("Hello, I'm %s (%d y.o.)\n", h.Name, h.Age)
}

func (h *Human) BirthDay() { h.Age++ }

func (h Human) Who() string { return "Human" }
