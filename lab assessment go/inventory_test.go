package main

import (
	"testing"
)

func TestAdd(t *testing.T) {
	inv := NewInventory()

	err := inv.Add(Item{ID: 1, Name: "Laptop", Quantity: 10, Price: 800.0})
	if err != nil {
		t.Fatalf("unexpected error adding item: %v", err)
	}

	err = inv.Add(Item{ID: 1, Name: "Duplicate", Quantity: 5, Price: 100.0})
	if err == nil {
		t.Fatal("expected error for duplicate ID, got nil")
	}

	err = inv.Add(Item{ID: 2, Name: "", Quantity: 5, Price: 50.0})
	if err == nil {
		t.Fatal("expected error for empty name, got nil")
	}

	err = inv.Add(Item{ID: 3, Name: "Mouse", Quantity: -1, Price: 20.0})
	if err == nil {
		t.Fatal("expected error for negative quantity, got nil")
	}

	err = inv.Add(Item{ID: 4, Name: "Keyboard", Quantity: 5, Price: -10.0})
	if err == nil {
		t.Fatal("expected error for negative price, got nil")
	}
}

func TestUpdate(t *testing.T) {
	inv := NewInventory()
	_ = inv.Add(Item{ID: 1, Name: "Pen", Quantity: 20, Price: 2.0})

	err := inv.Update(1, "Gel Pen", 25, 2.5)
	if err != nil {
		t.Fatalf("unexpected error updating item: %v", err)
	}

	items := inv.GetItems()
	if items[0].Name != "Gel Pen" || items[0].Quantity != 25 || items[0].Price != 2.5 {
		t.Fatal("item values were not updated correctly")
	}

	err = inv.Update(99, "NonExistent", 5, 10.0)
	if err == nil {
		t.Fatal("expected error updating non-existent item, got nil")
	}
}

func TestRemove(t *testing.T) {
	inv := NewInventory()
	_ = inv.Add(Item{ID: 1, Name: "Eraser", Quantity: 50, Price: 0.5})

	err := inv.Remove(1)
	if err != nil {
		t.Fatalf("unexpected error removing item: %v", err)
	}

	if len(inv.GetItems()) != 0 {
		t.Fatal("expected inventory to be empty")
	}

	err = inv.Remove(99)
	if err == nil {
		t.Fatal("expected error removing non-existent item, got nil")
	}
}

func TestComputeStats(t *testing.T) {
	inv := NewInventory()
	_ = inv.Add(Item{ID: 1, Name: "Notebook", Quantity: 5, Price: 20.0})
	_ = inv.Add(Item{ID: 2, Name: "Pencil", Quantity: 10, Price: 2.0})

	stats := inv.ComputeStats()

	if stats.TotalItems != 2 {
		t.Fatalf("expected 2 items, got %d", stats.TotalItems)
	}
	if stats.TotalQuantity != 15 {
		t.Fatalf("expected 15 total quantity, got %d", stats.TotalQuantity)
	}
	if stats.TotalValue != 120.0 {
		t.Fatalf("expected 120.0 total value, got %.2f", stats.TotalValue)
	}
}
