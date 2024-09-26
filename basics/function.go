package main

import "fmt"

func main() {
	greeting()
	greeting2(12)
	fmt.Println(sayHello(("Go!")))

	firstName, lastName := getFullName()
	fmt.Println(firstName)
	fmt.Println(lastName)

	firstName2, lastName2 := getFullName2()
	fmt.Println(firstName2)
	fmt.Println(lastName2)

	total1 := sumAll(10, 10, 10, 10, 10)
	fmt.Println(total1)

	slice := []int{20, 20, 20, 20, 20}
	total2 := sumAll(slice...)
	fmt.Println(total2)

	funcValue := sumAll
	fmt.Println(funcValue(10, 10, 20, 20, 20))

	sayHello2("GO", spamFilter)
	sayHello2("stupid", spamFilter)

	blackList := func(name string) bool {
		switch name {
		case "admin":
			return true
		case "root":
			return true
		default:
			return false
		}
	}

	registerUser("Odia", blackList)
	registerUser("admin", blackList)
	registerUser("root", blackList)
}

func greeting() {
	fmt.Println("Hello!")
}

func greeting2(hour int) {
	switch {
	case hour > 3 && hour < 12:
		fmt.Println("Good Morning!")
	case hour >= 12 && hour <= 18:
		fmt.Println("Noon!")
	case hour > 18:
		fmt.Println("Good Evening!")
	default:
		fmt.Println("Invalid hour!")
	}
}

func sayHello(name string) string {
	return "Hello " + name
}

func getFullName() (string, string) {
	return "Odia", "Pratama"
}

func getFullName2() (firstName, lastName string) {
	firstName = "Odia"
	lastName = "Pratama"
	return
}

func sumAll(numbers ...int) int {
	total := 0
	for _, value := range numbers {
		total += value
	}
	return total
}

//Filter Type
type Filter func(string) string

func sayHello2(name string, filter Filter) {
	fmt.Println("Hallo", filter(name))
}

func spamFilter(word string) string {
	if word == "stupid" {
		return "..."
	} else {
		return word
	}
}

//BlackList Type
type BlackList func(string) bool

func registerUser(name string, blackList BlackList) {
	if blackList(name) {
		fmt.Println("User blocked", name)
	} else {
		fmt.Println("You're in", name)
	}
}
