package main

import "fmt"

func main() {
	start()
	process(5)
	errorPanic(true)
	end()
}

func start() {
	fmt.Println("Start...")
}

func end() {
	fmt.Println("End...")
}

func process(number int) {
	for i := 1; i <= number; i++ {
		fmt.Println(i)
	}
}

func errorPanic(isPanic bool) {
	defer errorRecover()
	if isPanic {
		panic("Panic Error!")
	}
}

func errorRecover() {
	message := recover()
	if message != nil {
		fmt.Println("Recover message:", message)
	}
}
