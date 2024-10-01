package main

import (
	"fmt"
	"log"
	"os"
)

func readFile(filename string) ([]byte, error) {
	b, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("Error reading file: %v", err)
	}
	return b, nil
}

func wrapper() {
	println("=== Function Wrapper ===")
	b, err := readFile("news.txt")
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	fmt.Println(b)
	fmt.Println(string(b))
}
