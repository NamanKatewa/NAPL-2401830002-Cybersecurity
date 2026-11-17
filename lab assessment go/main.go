package main

import (
	"fmt"
)

func main() {
	inv := NewInventory()

	fmt.Println("=== 1. Adding Items ===")
	_ = inv.Add(Item{ID: 1, Name: "Laptop", Quantity: 5, Price: 800.0})
	_ = inv.Add(Item{ID: 2, Name: "Mouse", Quantity: 20, Price: 25.0})
	_ = inv.Add(Item{ID: 3, Name: "Keyboard", Quantity: 10, Price: 45.0})

	for _, item := range inv.GetItems() {
		fmt.Printf("ID: %d | Name: %-10s | Qty: %-3d | Price: $%.2f\n", item.ID, item.Name, item.Quantity, item.Price)
	}

	fmt.Println("\n=== 2. Updating Item 2 ===")
	err := inv.Update(2, "Wireless Mouse", 25, 30.0)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Updated item 2 successfully")
	}

	fmt.Println("\n=== 3. Removing Item 3 ===")
	err = inv.Remove(3)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Removed item 3 successfully")
	}

	fmt.Println("\n=== 4. Testing Input Validation ===")
	err = inv.Add(Item{ID: 1, Name: "Duplicate Laptop", Quantity: 2, Price: 800.0})
	if err != nil {
		fmt.Println("Duplicate ID check:", err)
	}

	err = inv.Add(Item{ID: 4, Name: "Faulty Item", Quantity: -5, Price: 10.0})
	if err != nil {
		fmt.Println("Negative quantity check:", err)
	}

	fmt.Println("\n=== 5. Current Inventory ===")
	for _, item := range inv.GetItems() {
		fmt.Printf("ID: %d | Name: %-15s | Qty: %-3d | Price: $%.2f\n", item.ID, item.Name, item.Quantity, item.Price)
	}

	fmt.Println("\n=== 6. Concurrent Summary Statistics ===")
	stats := inv.ComputeStats()
	fmt.Printf("Total Distinct Items : %d\n", stats.TotalItems)
	fmt.Printf("Total Quantity       : %d units\n", stats.TotalQuantity)
	fmt.Printf("Total Value          : $%.2f\n", stats.TotalValue)
}
