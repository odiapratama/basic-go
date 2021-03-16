package main

import (
	"fmt"
	"sort"
)

// User struct
type User struct {
	Name string
	Age  int
}

// UserSlice slice
type UserSlice []User

// Len UserSlice
func (value UserSlice) Len() int {
	return len(value)
}

// Less UserSlice
func (value UserSlice) Less(i, j int) bool {
	return value[i].Age < value[j].Age
}

func (value UserSlice) Swap(i, j int) {
	value[i], value[j] = value[j], value[i]
}

func main() {
	users := []User{
		{"Odia", 20},
		{"Pratama", 21},
		{"Budi", 23},
		{"Andi", 19},
	}

	fmt.Println(users)
	sort.Sort(UserSlice(users))
	fmt.Println(users)
}
