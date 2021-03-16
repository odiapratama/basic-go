package helper

import "fmt"

// SayGreeting (Public Function)
func SayGreeting(name string) {
	fmt.Println("Haloo", name)
	sayGoodBye(name)
}

// sayGoodBye (private function)
func sayGoodBye(name string) {
	fmt.Println("Goodbye ", name)
}
