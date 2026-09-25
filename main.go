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

		in1, err := strconv.ParseFloat(input1, 64)
		if err != nil {
			fmt.Println("\nMust be numeric number")
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
		case "+", "-", "*", "/":

		default:
			fmt.Println("\nInvalid option enter (+,-,*,/)")
			continue
		}

		var input2 string
		fmt.Printf("Enter value:")
		fmt.Scanln(&input2)

		in2, err := strconv.ParseFloat(input2, 64)
		if err != nil {
			fmt.Println("Must be number:", err)
			continue
		}

		switch operator {

		case "+":
			fmt.Printf("Result: %.2f", logics.Sum(in1, in2))
		case "-":
			fmt.Printf("Result: %.2f", logics.Sub(in1, in2))
		case "*":
			fmt.Printf("Result: %.2f", logics.Mul(in1, in2))
		case "/":
			result, err := logics.Div(in1, in2)
			if err != nil {
				fmt.Println(err)
				continue
			}
			fmt.Printf("\nResult: %.2f", result)
		}

	}
}
