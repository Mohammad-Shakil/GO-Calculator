package main

import "fmt"

func main() {

	fmt.Println("--- Calculator ---")

	for {

		var input1 string
		fmt.Printf("Enter value / Q to exit")
		fmt.Scanln(&input1)
		fmt.Printf("\n----Choose operator----")
		fmt.Println("1: +")
		fmt.Println("2: -")
		fmt.Println("3: *")
		fmt.Println("4: /")

	}
}
