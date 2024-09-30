package main

import "fmt"

func interfacePolymorph() {
	p := person{
		name: "Odia",
	}

	a := agent{
		person: p,
		ltk:    true,
	}

	fmt.Println(a)
	saySomething(p)
	saySomething(a)
}

type person struct {
	name string
}

type agent struct {
	person
	ltk bool
}

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

func (p person) speak() {
	fmt.Println("I am", p.name)
}

func (a agent) speak() {
	fmt.Println("I am", a.name, " - the agent")
}
