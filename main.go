package main

import (
	"errors"
	"fmt"
	"log"
)

// Defining custom errors
var (
	ErrNotImplemented = errors.New("not implemented")
	ErrTruckNotFound  = errors.New("truck not found")
)

// Truck interface defines the behavior (blueprint)
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
	e.cargo += 1
	e.battery -= 1
	return nil
}

func (e *ElectricTruck) UnloadCargo() error {
	e.cargo = 0
	e.battery -= 1
	return nil
}

// processTruck accepts the interface, not a concrete struct
func processTruck(t Truck) error {
	fmt.Printf("Processing truck\n")

	if err := t.LoadCargo(); err != nil {
		return fmt.Errorf("error loading cargo: %w", err)
	}

	if err := t.UnloadCargo(); err != nil {
		return fmt.Errorf("error unloading cargo: %w", err)
	}

	return nil
}

func main() {
	nt := &NormalTruck{id: "Normal-1"}
	et := &ElectricTruck{id: "Electric-1", battery: 100}

	// Both can be passed to processTruck because they implement the interface
	if err := processTruck(nt); err != nil {
		log.Fatalf("Error processing normal truck: %s", err)
	}

	if err := processTruck(et); err != nil {
		log.Fatalf("Error processing electric truck: %s", err)
	}

	fmt.Printf("Normal Truck Cargo: %d\n", nt.cargo)
	fmt.Printf("Electric Truck Battery: %.1f\n", et.battery)
}
