package main

import (
	"fmt"
	"strconv"
)

func main() {
	valueBool, _ := strconv.ParseBool("true")
	fmt.Println(valueBool)

	valueInt, _ := strconv.ParseInt("100", 10, 32)
	fmt.Println(valueInt)

	valueAtoi, _ := strconv.Atoi("100")
	fmt.Println(valueAtoi)
}
