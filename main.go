package main

import (
	"calcu/logics"
	"fmt"
)

func main() {

	fmt.Println("--- Calculator ---")

	for {

		var input1 string
		fmt.Printf("Q to exit\n Enter value:")
		fmt.Scanln(&input1)

		logics.Checker(input1)

		fmt.Printf("\n----Choose operator----")

		var operator string
		fmt.Println("\n1: +")
		fmt.Println("2: -")
		fmt.Println("3: *")
		fmt.Println("4: /")
		fmt.Scanln(&operator)

		var input2 string
		fmt.Printf("Enter value:")
		fmt.Scanln(&input2)

	}
}
