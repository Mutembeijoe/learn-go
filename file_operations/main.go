package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	rf, err := os.Open("./data.txt")

	if err != nil {
		fmt.Println("Failed to open file", err)
		os.Exit(1)
	}

	defer rf.Close()

	wf, err := os.Create("./write_file.txt")

	if err != nil {
		fmt.Println("Failed to Create file", err)
		os.Exit(1)
	}

	defer wf.Close()

	n, err := io.Copy(wf, rf)

	if err != nil {
		fmt.Println("Failed to copy file data")
		os.Exit(1)
	}

	fmt.Printf("Copied %d from %s to %s ", n, rf.Name(), wf.Name())

}
