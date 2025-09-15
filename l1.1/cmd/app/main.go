package main

import (
	"fmt"

	"l1.1/internal/domain"
)

func main() {
	a := domain.Action{
		Human: domain.Human{Name: "Sergey", Age: 20},
		Role:  "DevOps",
	}

	a.Greet()
	a.BirthDay()
	a.Greet()
	fmt.Println(a.Who())
	fmt.Println(a.Human.Who())

	a.LevelUp()
}
