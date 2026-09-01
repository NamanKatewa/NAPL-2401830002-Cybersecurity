package main

import (
	"fmt"
	"strconv"
)

func test() {
	// fmt.Println("Hello! Naman")
	// fmt.Println()
	// var a int = 10
	// var b int = 20
	// var c float32 = 10.8
	// var d float32 = 20.9

	// fmt.Printf("Integer Addition: %d + %d = %d\n", a, b, a+b)
	// fmt.Printf("Integer Subtraction: %d - %d = %d\n", a, b, a-b)
	// fmt.Printf("Integer Multiplication: %d X %d = %d\n\n", a, b, a*b)
	// fmt.Printf("FLoat Addition: %f + %f = %f\n", c, d, c+d)
	// fmt.Printf("Float Subtraction: %f - %f = %f\n", c, d, c+d)
	// fmt.Printf("Float Multiplication: %f X %f = %f\n", c, d, c+d)

	var choice string
	for {
		fmt.Println("\n--- Calculator ---")
		fmt.Println("1. Integer Operations (Add, Sub, Mul)")
		fmt.Println("2. Floating-Point Operations (Add, Sub, Mul)")
		fmt.Println("3. Exit")
		fmt.Print("Enter choice (1-3): ")

		fmt.Scan(&choice)

		if choice == "3" {
			fmt.Println("Exiting...")
			break
		} else if choice == "1" {
			var input1, input2 string
			var a, b int
			var err error

			for {
				fmt.Print("Enter first integer: ")
				fmt.Scan(&input1)
				a, err = strconv.Atoi(input1)
				if err == nil {
					break
				}
				fmt.Println("Invalid input. Please enter a valid integer.")
			}

			for {
				fmt.Print("Enter second integer: ")
				fmt.Scan(&input2)
				b, err = strconv.Atoi(input2)
				if err == nil {
					break
				}
				fmt.Println("Invalid input. Please enter a valid integer.")
			}

			fmt.Printf("Integer Addition: %d + %d = %d\n", a, b, a+b)
			fmt.Printf("Integer Subtraction: %d - %d = %d\n", a, b, a-b)
			fmt.Printf("Integer Multiplication: %d * %d = %d\n", a, b, a*b)

		} else if choice == "2" {
			var input1, input2 string
			var c, d float64
			var err error

			for {
				fmt.Print("Enter first float: ")
				fmt.Scan(&input1)
				c, err = strconv.ParseFloat(input1, 64)
				if err == nil {
					break
				}
				fmt.Println("Invalid input. Please enter a valid floating-point number.")
			}

			for {
				fmt.Print("Enter second float: ")
				fmt.Scan(&input2)
				d, err = strconv.ParseFloat(input2, 64)
				if err == nil {
					break
				}
				fmt.Println("Invalid input. Please enter a valid floating-point number.")
			}

			fmt.Printf("Float Addition: %f + %f = %f\n", c, d, c+d)
			fmt.Printf("Float Subtraction: %f - %f = %f\n", c, d, c-d)
			fmt.Printf("Float Multiplication: %f * %f = %f\n", c, d, c*d)
		} else {
			fmt.Println("Invalid choice. Please enter 1, 2, or 3.")
		}
	}
}
