package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Person struct {
	Name   string
	Age    int
	Job    string
	Salary float64
}

var reader = bufio.NewReader(os.Stdin)

func (p *Person) Read() {
	fmt.Print("Enter Name: ")
	name, _ := reader.ReadString('\n')
	p.Name = strings.TrimSpace(name)

	for {
		fmt.Print("Enter Age: ")
		ageStr, _ := reader.ReadString('\n')
		age, err := strconv.Atoi(strings.TrimSpace(ageStr))
		if err == nil && age >= 0 && age <= 150 {
			p.Age = age
			break
		}
		fmt.Println("Invalid age! Must be more than 0 and below 150. Try again.")
	}

	fmt.Print("Enter Job: ")
	job, _ := reader.ReadString('\n')
	p.Job = strings.TrimSpace(job)

	for {
		fmt.Print("Enter Salary: ")
		salaryStr, _ := reader.ReadString('\n')
		salary, err := strconv.ParseFloat(strings.TrimSpace(salaryStr), 64)
		if err == nil && salary >= 0 {
			p.Salary = salary
			break
		}
		fmt.Println("Invalid salary! Cannot be negative. Try again.")
	}
}

func (p Person) Print() {
	fmt.Printf("Name:   %s\n", p.Name)
	fmt.Printf("Age:    %d\n", p.Age)
	fmt.Printf("Job:    %s\n", p.Job)
	fmt.Printf("Salary: $%.2f\n", p.Salary)
	fmt.Println(strings.Repeat("-", 25))
}

func modifyValue(val *int) {
	*val = *val + 100
}

func demonstratePointers() {
	fmt.Println("\n========== Pointer Demonstration ==========")

	var myVar int = 42
	fmt.Printf("Original variable value: %d\n", myVar)
	fmt.Printf("Variable address (&myVar): %p\n", &myVar)

	var myPtr *int = &myVar
	fmt.Printf("Value accessed using pointer (*myPtr): %d\n", *myPtr)

	fmt.Printf("\nBefore modifyValue: myVar = %d\n", myVar)
	modifyValue(&myVar)
	fmt.Printf("After modifyValue: myVar = %d\n", myVar)

	personPtr := new(Person)
	fmt.Printf("\nNewly allocated Person struct (using new): %+v\n", *personPtr)

	// Modify fields through the pointer (implicitly and explicitly)
	personPtr.Name = "Alice"
	(*personPtr).Age = 30
	personPtr.Job = "Engineer"
	personPtr.Salary = 90000.0

	fmt.Printf("Modified Person struct accessed through pointer: %+v\n", *personPtr)
}

func main() {
	var persons []Person

	for {
		var p Person
		fmt.Printf("Enter details for Person %d\n", len(persons)+1)
		p.Read()
		persons = append(persons, p)

		fmt.Print("Do you want to add another person? (y/n): ")
		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(strings.ToLower(choice))

		if choice != "y" && choice != "yes" {
			break
		}
		fmt.Println()
	}

	fmt.Println("\n========== Output ==========")
	for _, p := range persons {
		p.Print()
	}

	demonstratePointers()
}
