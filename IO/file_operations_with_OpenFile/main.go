package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	rf, err := os.OpenFile("./write_only.txt", os.O_WRONLY, 0666)

	if err != nil {
		fmt.Println("failed to Open file", err)
		os.Exit(1)
	}

	defer rf.Close()

	newReader := strings.NewReader("Hello World, This is a strings reader")

	n, err := io.Copy(rf, newReader)

	if err != nil {
		fmt.Println("Failed to copy data", err)
		os.Exit(1)
	}

	fmt.Printf("Copied %d bytes from %s to StdOut", n, rf.Name())

}
