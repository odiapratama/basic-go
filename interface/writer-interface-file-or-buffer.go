package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
)

func (p person) writeOut(w io.Writer) {
	w.Write([]byte(p.name))
}

func interfaceFileOrBuffer() {
	fmt.Println("=== Interface File or Buffer ===")

	p := person{
		name: "Odia",
	}

	f, err := os.Create("output.txt")
	if err != nil {
		log.Fatalf("error %s", err)
	}
	defer f.Close()

	var b bytes.Buffer

	p.writeOut(f)
	p.writeOut(&b)
	fmt.Println(b.String())
}
