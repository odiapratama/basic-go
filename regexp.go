package main

import (
	"fmt"
	"regexp"
)

func main() {
	regex := regexp.MustCompile("a([a-z])i")

	fmt.Println(regex.MatchString("adi"))

	fmt.Println(regex.FindAllString("ari asi aWi dio kia aji", -1))

}
