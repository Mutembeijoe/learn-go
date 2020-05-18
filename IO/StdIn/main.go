package main

import "fmt"

var choice int

func main() {
	fmt.Println("What is a Square?")
	fmt.Print("Enter 1=Quadrilateral 2=Rectahedron :")

	n, err := fmt.Scanf("%d", &choice)

	if n != 1 || err != nil {
		fmt.Print("Follow the Instructions!")
		return
	}

	if choice == 1 {
		fmt.Println("Correct!")
	} else {
		fmt.Println("Failed, Google it")
	}
}
