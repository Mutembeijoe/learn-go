package main

import (
	"fmt"
	"os"
)

func main() {
	strSlice := []string{
		"Hello World",
		"Data IO operatiosn in go",
		"cool, isn't it?",
	}

	byteSlice := [][]byte{
		[]byte("Hello World"),
		[]byte("Goland Data IO"),
	}

	sf, err := os.Create("./slice.txt")

	if err != nil {
		fmt.Println("Error creating file", err)
		os.Exit(1)
	}

	bf, err := os.Create("./byte.txt")

	if err != nil {
		fmt.Println("Error creating file", err)
		os.Exit(1)
	}

	for _, row := range strSlice {
		// fmt.Fprintf(wf, " %s \n", row)
		sf.WriteString(row)
	}

	for _, row := range byteSlice {
		bf.Write(row)
	}
}
