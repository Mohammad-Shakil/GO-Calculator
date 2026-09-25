package main

import (
	"calcu/logics"
	"fmt"
	"strconv"
)

func main() {

	fmt.Println("--- Calculator ---")

	for {

		var input1 string
		fmt.Printf("\nQ to exit\nEnter value:")
		fmt.Scanln(&input1)

		switch input1 {
		case "Q", "q":
			return
		}

		in1, err := strconv.Atoi(input1)
		if err != nil {
			fmt.Println("\nMust be numeric number:", err)
			continue
		}

		fmt.Printf("\n----Choose operator----")

		var operator string
		fmt.Println("\n1: +")
		fmt.Println("2: -")
		fmt.Println("3: *")
		fmt.Println("4: /")
		fmt.Scanln(&operator)

		switch operator {
		case "1", "2", "3", "4":

		default:
			fmt.Println("\nInvalid option:", operator)
			continue
		}

		var input2 string
		fmt.Printf("Enter value:")
		fmt.Scanln(&input2)

		in2, err := strconv.Atoi(input2)
		if err != nil {
			fmt.Println("Must be number:", err)
			continue
		}

		switch operator {

		case "1":
			fmt.Println("Result:", logics.Add(in1, in2))
		case "2":
			fmt.Println("Result:", logics.Sub(in1, in2))
		case "3":
			fmt.Println("Result:", logics.Mul(in1, in2))
		case "4":
			result, err := logics.Div(in1, in2)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println("Result:", result)
		}

	}
}
