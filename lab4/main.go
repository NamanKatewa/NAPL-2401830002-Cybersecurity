package main

import (
	"fmt"
	"sync"
	"time"
)

// TaskResult defines the structure for holding results from our concurrent tasks
type TaskResult struct {
	Name string
	Data []int
}

// computeSquares computes the squares of numbers from 1 to n
func computeSquares(n int, ch chan<- TaskResult, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("[Start] Goroutine: Computing Squares")
	
	var data []int
	for i := 1; i <= n; i++ {
		data = append(data, i*i)
		// Small sleep to simulate workload and allow goroutines to interleave
		time.Sleep(10 * time.Millisecond)
	}
	
	ch <- TaskResult{Name: "Squares", Data: data}
	fmt.Println("[Complete] Goroutine: Computing Squares")
}

// computeCubes computes the cubes of numbers from 1 to n
func computeCubes(n int, ch chan<- TaskResult, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("[Start] Goroutine: Computing Cubes")
	
	var data []int
	for i := 1; i <= n; i++ {
		data = append(data, i*i*i)
		time.Sleep(15 * time.Millisecond)
	}
	
	ch <- TaskResult{Name: "Cubes", Data: data}
	fmt.Println("[Complete] Goroutine: Computing Cubes")
}

// computeFibonacci computes the first n Fibonacci numbers
func computeFibonacci(n int, ch chan<- TaskResult, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("[Start] Goroutine: Computing Fibonacci")
	
	var data []int
	a, b := 0, 1
	for i := 0; i < n; i++ {
		data = append(data, a)
		a, b = b, a+b
		time.Sleep(12 * time.Millisecond)
	}
	
	ch <- TaskResult{Name: "Fibonacci", Data: data}
	fmt.Println("[Complete] Goroutine: Computing Fibonacci")
}

func main() {
	fmt.Println("========== Concurrency Demonstration ==========")

	// Buffered channel to collect exactly 3 results
	ch := make(chan TaskResult, 3)
	var wg sync.WaitGroup
	
	// Number of elements to compute for each task
	n := 10

	// Launch 3 independent goroutines
	wg.Add(3)
	go computeSquares(n, ch, &wg)
	go computeCubes(n, ch, &wg)
	go computeFibonacci(n, ch, &wg)

	// Wait for all goroutines to finish before closing the channel
	wg.Wait()
	close(ch)

	fmt.Println("\n========== Collected Results ==========")
	// Read and print results from the channel in the main goroutine
	for result := range ch {
		fmt.Printf("%-10s: %v\n", result.Name, result.Data)
	}
}
