package main

import (
	"fmt"
	"reflect"
)

type Sample struct {
	Name    string `required:"true"`
	Address string `required:"false"`
}

func isValid(data interface{}) bool {
	t := reflect.TypeOf(data)
	for i := 0; i < t.NumField(); i++ {
		if t.Field(i).Tag.Get("required") == "true" {
			return reflect.ValueOf(data).Field(i).Interface() != ""
		}
	}
	return true
}

func main() {
	sample := Sample{"James", "Bandung"}
	sampleType := reflect.TypeOf(sample)

	fmt.Println(sampleType.NumField())
	for i := 0; i < sampleType.NumField(); i++ {
		fmt.Println(sampleType.Field(i).Name, sampleType.Field(i).Tag.Get("required"))
	}

	fmt.Println(isValid(sample))

	sample2 := Sample{"", ""}
	fmt.Println(isValid(sample2))
}
