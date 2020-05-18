package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	rows := []byte("Hello World from Golang")

	wf, err := os.Create("./write.txt")

	if err != nil {
		fmt.Println("Error when creating file", err)
		os.Exit(1)
	}

	writer := bufio.NewWriter(wf)

	for _, row := range rows {
		err := writer.WriteByte(row)
		if err != nil {
			fmt.Println("Failed to write bytes", err)
		}
	}

	writer.Flush()

}
