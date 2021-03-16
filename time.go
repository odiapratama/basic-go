package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now)

	utc := time.Date(2020, time.December, 5, 5, 5, 5, 5, time.UTC)
	fmt.Println(utc)

	parse, err := time.Parse(time.RFC3339, "2020-10-10T15:04:05Z07:00")
	if err == nil {
		fmt.Println(parse)
	} else {
		fmt.Println(err)
	}
}
