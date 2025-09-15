package domain

import "fmt"

type Action struct {
	Human
	Role string
}

func (a Action) Who() string { return "Action/" + a.Role }

func (a Action) LevelUp() {
	a.BirthDay()
	a.Human.Greet()
	fmt.Println(a.Who())
}

type WhoGreeter interface {
	Greet()
	Who() string
}
