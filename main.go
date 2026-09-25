package main

import (
	"calcu/logics"
	"fmt"
	"strconv"
)

func main() {

	fmt.Println("--Calculator--")
	for {

		var input1 string
		fmt.Print("\npress Q to exit \nEnter number :")
		fmt.Scanln(&input1)
		switch input1 {
		case "Q", "q":
			return
		}
		in1, err := strconv.ParseFloat(input1, 64)
		if err != nil {
			fmt.Println("\nMust be number")
			continue
		}

		var operator string
		fmt.Println("\nChoose operator")
		fmt.Println("\n+")
		fmt.Println("-")
		fmt.Println("*")
		fmt.Println("/")
		fmt.Print("Enter operator:")
		fmt.Scanln(&operator)
		switch operator {
		case "+", "-", "*", "/":
		default:
			fmt.Println("\nInvalid operator")
			continue
		}

		var input2 string
		fmt.Print("Enter number:")
		fmt.Scanln(&input2)

		in2, err := strconv.ParseFloat(input2, 64)
		if err != nil {
			fmt.Println("Must be number")
		}

		switch operator {
		case "+":
			fmt.Printf("\nResult:%.2f", logics.Add(in1, in2))

		case "-":
			fmt.Printf("\nResult:%.2f", logics.Sub(in1, in2))
		case "*":
			fmt.Printf("\nResult:%.2f", logics.Mul(in1, in2))
		case "/":
			res, err := logics.Div(in1, in2)
			if err != nil {
				fmt.Print(err)
				continue
			}
			fmt.Printf("Result:%.2f", res)
		}

	}
}
