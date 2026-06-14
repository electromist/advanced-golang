package main

import (
	"errors"
	"fmt"
	"log"
)

var (
	ErrNotImplemented = errors.New("not implemented")
	ErrTruckNotFound  = errors.New("truck not found")
)

// Truck interface defines the behavior for any type of truck
type Truck interface {
	LoadCargo() error
	UnloadCargo() error
}

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

type ElectricTruck struct {
	id      string
	cargo   int
	battery int
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

// processTruck now accepts the Truck interface, not a concrete struct
func processTruck(t Truck) error {
	if err := t.LoadCargo(); err != nil {
		return fmt.Errorf("error loading cargo: %w", err)
	}
	return t.UnloadCargo()
}

func main() {
	normalTruck := &NormalTruck{id: "Truck-1"}
	electricTruck := &ElectricTruck{id: "Truck-2", battery: 100}

	// Processing trucks through the interface
	trucks := []Truck{normalTruck, electricTruck}

	for _, t := range trucks {
		if err := processTruck(t); err != nil {
			log.Fatalf("Error processing truck: %v", err)
		}
		fmt.Printf("Processed: %+v\\n", t)
	}
}
