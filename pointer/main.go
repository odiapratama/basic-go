package main

import "fmt"

func main() {
	pointerMethod()
	pointer()

	fmt.Println("\n=====================================\n")

	x := 42
	y := &x
	z := *y

	fmt.Printf("%v\t\t%T\n", x, x)
	fmt.Printf("%v\t%T\n", y, y)
	fmt.Printf("%v\t\t%T\n", z, z)

	fmt.Println("\n=====================================\n")

	*y = 43
	z = 44

	fmt.Printf("%v\t\t%T\n", x, x)
	fmt.Printf("%v\t%T\n", y, y)
	fmt.Printf("%v\t\t%T\n", z, z)

	fmt.Println("\n=====================================\n")
	a := Int(z).addValue(1)
	fmt.Printf("%v\t%T\n", a, a)
	fmt.Printf("%v\t\t%T\n", z, z)
	addValue2(&z)
	fmt.Printf("%v\t%T\n", a, a)
	fmt.Printf("%v\t\t%T\n", z, z)

	fmt.Println("\n=====================================\n")
	d := animal{"Bobby"}
	d.run()
	fmt.Println(d)
	changeName2(d)
	d.walk()
	fmt.Println(d)
	changeName(&d)
	fmt.Println(d)
	d.run()
}

type Int int

func (v Int) addValue(n int) int {
	return int(v) + n
}

func addValue2(n *int) {
	*n += 1
}

func changeName(d *animal) {
	d.name = "Bobby Jr."
}

func changeName2(d animal) {
	d.name = "Bobby Jr."
}

func (a animal) walk() {
	fmt.Println(a.name, "Walking...")
}

func (a animal) run() {
	fmt.Println(a.name, "Running...")
}

type animal struct {
	name string
}

type animalActivity interface {
	walk()
	run()
}

func animalRun(a animalActivity) {
	a.run()
}

func animalWalk(a animalActivity) {
	a.walk()
}
