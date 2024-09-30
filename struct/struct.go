package main

import "fmt"

// Employee struct
type Employee struct {
	ID, Name, Title string
	Age             int
}

// Director struct
type Director struct {
	Name string
}

func (employee Employee) sayHello() {
	fmt.Println("Hi", employee.Name)
}

//Greeting Interface
type Greeting interface {
	GetGreeting() string
}

// GetGreeting employee
func (employee Employee) GetGreeting() {
	fmt.Println("Hallo employee", employee.Name)
}

// GetGreeting director
func (director Director) GetGreeting() {
	fmt.Println("Hello director", director.Name)
}

func main() {
	var employee1 Employee
	employee1.ID = "123123142"
	employee1.Name = "James"
	employee1.Title = "IoT Enginerr"
	employee1.Age = 20

	var employee2 = Employee{
		ID:    "32u598239",
		Name:  "George",
		Title: "Android Developer",
		Age:   22,
	}

	var employee3 = Employee{"382749283749", "Ahmad", "Backend Engineer", 24}

	fmt.Println(employee1)
	fmt.Println(employee2)
	fmt.Println(employee3)

	employee1.sayHello()
	employee2.sayHello()
	employee3.sayHello()

	employee1.GetGreeting()

	director := Director{"Michael"}
	director.GetGreeting()
}
