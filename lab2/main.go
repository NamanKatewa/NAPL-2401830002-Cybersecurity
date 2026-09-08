package main

import (
	"bufio"
	"fmt"
	"lab2/utils"
	"os"
	"strconv"
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

	demonstrateCollections()
}

func demonstrateCollections() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("--- Slice Operations ---")
	students := []string{"Alice", "Bob", "Charlie"}
	fmt.Printf("Initial slice: %v\n", students)

	// Add an element
	fmt.Print("\nEnter a student name to add: ")
	scanner.Scan()
	newStudent := scanner.Text()
	students = append(students, newStudent)
	fmt.Printf("After adding '%s': %v\n", newStudent, students)

	// Update an element
	fmt.Printf("\nEnter an index to update (0 to %d): ", len(students)-1)
	scanner.Scan()
	if idx, err := strconv.Atoi(scanner.Text()); err == nil && idx >= 0 && idx < len(students) {
		fmt.Print("Enter the new name: ")
		scanner.Scan()
		updatedName := scanner.Text()
		students[idx] = updatedName
		fmt.Printf("After updating index %d: %v\n", idx, students)
	} else {
		fmt.Println("Invalid index! Skipping update.")
	}

	// Remove an element
	fmt.Printf("\nEnter an index to remove (0 to %d): ", len(students)-1)
	scanner.Scan()
	if idx, err := strconv.Atoi(scanner.Text()); err == nil && idx >= 0 && idx < len(students) {
		students = append(students[:idx], students[idx+1:]...)
		fmt.Printf("After removing index %d: %v\n", idx, students)
	} else {
		fmt.Println("Invalid index! Skipping removal.")
	}

	fmt.Println("\n--- Map Operations ---")
	marks := map[string]int{
		"Math":    90,
		"Physics": 85,
	}
	fmt.Printf("Initial map: %v\n", marks)

	// Insert / Update
	fmt.Print("\nEnter a subject to add or update: ")
	scanner.Scan()
	subject := scanner.Text()

	fmt.Printf("Enter the score for %s: ", subject)
	scanner.Scan()
	if score, err := strconv.Atoi(scanner.Text()); err == nil {
		marks[subject] = score
		fmt.Printf("After updating '%s': %v\n", subject, marks)
	} else {
		fmt.Println("Invalid score! Skipping insertion/update.")
	}

	// Lookup
	fmt.Print("\nEnter a subject to lookup: ")
	scanner.Scan()
	lookupSubject := scanner.Text()
	if score, exists := marks[lookupSubject]; exists {
		fmt.Printf("Lookup '%s': found with score %d\n", lookupSubject, score)
	} else {
		fmt.Printf("Lookup '%s': not found\n", lookupSubject)
	}

	// Delete
	fmt.Print("\nEnter a subject to delete: ")
	scanner.Scan()
	subjectToDelete := scanner.Text()
	delete(marks, subjectToDelete)
	fmt.Printf("After deleting '%s': %v\n", subjectToDelete, marks)
}
