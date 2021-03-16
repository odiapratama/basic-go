package database

import "fmt"

var connection string

func init() {
	fmt.Println("Initialization database...")
	connection = "Database connected!"

}

// GetConnection function
func GetConnection() string {
	return connection
}
