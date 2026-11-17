package main

import (
	"errors"
)

type Item struct {
	ID       int
	Name     string
	Quantity int
	Price    float64
}

type SummaryStats struct {
	TotalItems    int
	TotalQuantity int
	TotalValue    float64
}

type Inventory struct {
	items []Item
}

func NewInventory() *Inventory {
	return &Inventory{
		items: []Item{},
	}
}

func (inv *Inventory) Add(item Item) error {
	if item.ID <= 0 {
		return errors.New("item ID must be positive")
	}
	if item.Name == "" {
		return errors.New("item name cannot be empty")
	}
	if item.Quantity < 0 {
		return errors.New("quantity cannot be negative")
	}
	if item.Price < 0 {
		return errors.New("price cannot be negative")
	}

	for _, it := range inv.items {
		if it.ID == item.ID {
			return errors.New("item with this ID already exists")
		}
	}

	inv.items = append(inv.items, item)
	return nil
}

func (inv *Inventory) Update(id int, name string, quantity int, price float64) error {
	if name == "" {
		return errors.New("item name cannot be empty")
	}
	if quantity < 0 {
		return errors.New("quantity cannot be negative")
	}
	if price < 0 {
		return errors.New("price cannot be negative")
	}

	for i, it := range inv.items {
		if it.ID == id {
			inv.items[i].Name = name
			inv.items[i].Quantity = quantity
			inv.items[i].Price = price
			return nil
		}
	}

	return errors.New("item not found")
}

func (inv *Inventory) Remove(id int) error {
	for i, it := range inv.items {
		if it.ID == id {
			inv.items = append(inv.items[:i], inv.items[i+1:]...)
			return nil
		}
	}
	return errors.New("item not found")
}

func (inv *Inventory) GetItems() []Item {
	return inv.items
}

func (inv *Inventory) ComputeStats() SummaryStats {
	ch := make(chan SummaryStats)

	go func() {
		var stats SummaryStats
		stats.TotalItems = len(inv.items)

		for _, it := range inv.items {
			stats.TotalQuantity += it.Quantity
			stats.TotalValue += float64(it.Quantity) * it.Price
		}

		ch <- stats
	}()

	return <-ch
}
