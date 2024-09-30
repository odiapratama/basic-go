package main

import (
	"log"
	"os"
)

func writeInterfaceFile() {
	f, err := os.Create("output.txt")
	if err != nil {
		log.Fatalf("error %s", err)
	}
	defer f.Close()

	s := []byte("Hello, World!")

	_, err = f.Write(s)
	if err != nil {
		log.Fatalf("error %s", err)
	}
}
