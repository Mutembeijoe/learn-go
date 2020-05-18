package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
)

func main() {
	var books bytes.Buffer

	books.WriteString("Hello there, \n")
	books.WriteString("This in-memoty IO in golang. \n")
	books.WriteString("Go is pretty cool, Isn't it ?")

	file, err := os.Create("./dump.txt")

	if err != nil {
		log.Fatalln("Failed to create dump file", err)
	}

	n, err := books.WriteTo(file)

	if err != nil {
		log.Fatalln("Write to dump file failed", err)
	}

	fmt.Printf("Copied %d bytes to %s file", n, file.Name())
}
