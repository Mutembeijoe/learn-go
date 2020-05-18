package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {

	rf, err := os.Open("./read.txt")

	if err != nil {
		fmt.Println("Failed to open file", err)
		os.Exit(1)
	}

	reader := bufio.NewReader(rf)

	for {
		line, err := reader.ReadString('\n')

		if err != nil {
			if err == io.EOF {
				fmt.Println(err)
				break
			}
			fmt.Println("Error Reading", err)
			return
		}
		fmt.Print(line)
	}

}
