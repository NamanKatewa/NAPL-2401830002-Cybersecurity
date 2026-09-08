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
	fmt.Printf("%d to the power of %d: %d\n\n", base, exponent, utils.Power(base, exponent))

	collections()
}

func collections() {
	fmt.Println("--- Slice Operations ---")

	students := []string{"Alice", "Bob", "Charlie"}
	fmt.Printf("Initial slice: %v\n", students)

	students = append(students, "David")
	fmt.Printf("After adding 'David': %v\n", students)

	if len(students) > 1 {
		students[1] = "Bob (Updated)"
		fmt.Printf("After updating index 1: %v\n", students)
	}

	indexToRemove := 2
	if indexToRemove < len(students) {
		students = append(students[:indexToRemove], students[indexToRemove+1:]...)
		fmt.Printf("After removing index 2: %v\n", students)
	}

	fmt.Println()

	fmt.Println("--- Map Operations ---")

	marks := map[string]int{
		"Math":    90,
		"Physics": 85,
	}
	fmt.Printf("Initial map: %v\n", marks)

	marks["Chemistry"] = 92
	fmt.Printf("After inserting 'Chemistry': %v\n", marks)

	marks["Physics"] = 88
	fmt.Printf("After updating 'Physics' to 88: %v\n", marks)

	subject := "Math"
	if score, exists := marks[subject]; exists {
		fmt.Printf("Lookup '%s': found with score %d\n", subject, score)
	} else {
		fmt.Printf("Lookup '%s': not found\n", subject)
	}

	delete(marks, "Math")
	fmt.Printf("After deleting 'Math': %v\n", marks)
}
