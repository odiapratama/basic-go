package main

import (
	"fmt"
	"log"
	"os"
)

func writeInterfaceFile() {
	fmt.Println("=== Write Interface File ===")
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
