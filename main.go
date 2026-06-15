package main

import (
	"errors"
	"fmt"
	"log"
	"sync"
	"time"
)

// Custom errors
var (
	ErrNotImplemented = errors.New("not implemented")
	ErrTruckNotFound  = errors.New("truck not found")
)

// Truck interface defines the behavior
type Truck interface {
	LoadCargo() error
	UnloadCargo() error
}

// NormalTruck struct implementation
type NormalTruck struct {
	id    string
	cargo int
}

func (n *NormalTruck) LoadCargo() error {
	time.Sleep(1 * time.Second) // Simulating heavy work
	n.cargo += 1
	return nil
}

func (n *NormalTruck) UnloadCargo() error {
	n.cargo = 0
	return nil
}

// ElectricTruck struct implementation
type ElectricTruck struct {
	id      string
	cargo   int
	battery float64
}

func (e *ElectricTruck) LoadCargo() error {
	time.Sleep(1 * time.Second) // Simulating heavy work
	e.cargo += 1
	e.battery -= 1
	return nil
}

func (e *ElectricTruck) UnloadCargo() error {
	e.cargo = 0
	e.battery -= 1
	return nil
}

// processTruck single truck ko execute karta hai
func processTruck(t Truck) error {
	fmt.Printf("Started processing truck: %v\n", t)

	if err := t.LoadCargo(); err != nil {
		return fmt.Errorf("error loading cargo: %w", err)
	}

	if err := t.UnloadCargo(); err != nil {
		return fmt.Errorf("error unloading cargo: %w", err)
	}

	fmt.Printf("Finished processing truck: %v\n", t)
	return nil
}

// processFleet handles concurrent processing of multiple trucks
func processFleet(trucks []Truck) error {
	var wg sync.WaitGroup

	for _, t := range trucks {
		wg.Add(1) // Counter ko 1 se badhao har ek truck ke liye

		// Goroutine start karne ke liye 'go' keyword lagaya aur closure pass kiya
		go func(truck Truck) {
			if err := processTruck(truck); err != nil {
				log.Println(err) // Goroutine ke andar error log kar rahe hain
			}
			wg.Done() // Goroutine khatam hote hi counter ko 1 kam karo
		}(t) // Variable pinning issue se bachne ke liye 't' ko parameter pass kiya
	}

	wg.Wait() // Jab tak counter 0 nahi hota, tab tak main execution ko roko
	return nil
}

func main() {
	fleet := []Truck{
		&NormalTruck{id: "NT1", cargo: 0},
		&ElectricTruck{id: "ET1", cargo: 0, battery: 100},
		&NormalTruck{id: "NT2", cargo: 0},
		&ElectricTruck{id: "ET2", cargo: 0, battery: 100},
	}

	// Fleet ko concurrently process kar rahe hain
	if err := processFleet(fleet); err != nil {
		fmt.Printf("Error processing fleet: %v\n", err)
		return
	}

	fmt.Println("All trucks processed successfully!")
}
