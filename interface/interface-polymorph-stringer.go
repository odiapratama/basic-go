package main

import (
	"fmt"
)

type book struct {
	title string
}

func (b book) String() string {
	return fmt.Sprint("Book: ", b.title)
}

type count int

func (c count) String() string {
	// return fmt.Sprint("Count: ", strconv.Itoa(int(c)))
	return fmt.Sprint("Count: ", int(c))
}

func logInfo(i fmt.Stringer) {
	fmt.Println("Logging:", i.String())
}

func interfaceStringer() {
	b := book{
		title: "The Art of War",
	}

	var c count = 10

	fmt.Println(b)
	fmt.Println(c)

	logInfo(b)
	logInfo(c)
}
