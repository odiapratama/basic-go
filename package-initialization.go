package main

import (
	"explore/database"
	"fmt"
)

func main() {
	fmt.Println(database.GetConnection())
}
