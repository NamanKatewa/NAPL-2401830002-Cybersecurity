package main

import (
	"fmt"
	"lab2/utils"
)

func main() {
	fmt.Println("--- String Manipulation ---")
	
	str := "Cybersecurity"
	fmt.Printf("Original String: %s\n", str)
	
	reversedStr := utils.Reverse(str)
	fmt.Printf("Reversed: %s\n", reversedStr)
	
	vowelCount := utils.CountVowels(str)
	fmt.Printf("Vowels: %d\n\n", vowelCount)
	
	fmt.Println("--- Mathematical Utilities ---")
	
	num := 5
	fmt.Printf("Factorial of %d: %d\n", num, utils.Factorial(num))
	
	base := 2
	exponent := 8
	fmt.Printf("%d to the power of %d: %d\n", base, exponent, utils.Power(base, exponent))
}
