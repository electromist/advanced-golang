package main

import (
	"errors"
	"fmt"
	"log"
)

var (
	ErrNotImplemented = errors.New("Not Implemented")
	ErrTruckNotFound  = errors.New("Truck Not Found")
)

type Truck struct {
	id string
}

func (t *Truck) LoadCargo() error {
	return ErrTruckNotFound
}

// processTruck handles the loading
// error bas ek interface hai, aur jab humlog iska definition dekhte hai toh udhar bhi yahi dikhta hai
//
//	type error interface {
//		Error() string
//	}
func processTruck(truck Truck) error {
	fmt.Printf("Processing truck: %s\n", truck.id)
	// return errors.New("Not Implemented") // This is how easy its to create custom errors in go.
	// return ErrNotImplemented -> We can use this here now multiple times
	// We can use errors.Is(err, SomeDefinedError) -> This is for case-matching using if or switch case and 2nd way

	if err := truck.LoadCargo(); err != nil {
		return fmt.Errorf("Error loading Cargo: %w", err)
	}
	return ErrNotImplemented
}

// in golang we handle errors in a declaritive way
// we explicity define it
func main() {
	trucks := []Truck{
		{id: "Truck-1"},
		{id: "Truck-2"},
		{id: "Truck-3"},
	}

	for _, truck := range trucks {
		fmt.Printf("Truck %s arrived.\n", truck.id)
		// 1st Way
		// err := processTruck(truck)
		// if err != nil {
		// 	log.Fatalf("Error processing truck %s:\n", err)
		// }

		// 2nd Way -> More Cleaner
		if err := processTruck(truck); err != nil {
			// if errors.Is(err, ErrNotImplemented) {
			// 	// we do something for this specefic error
			// }
			// or we can use switch case
			log.Fatalf("Error processing truck %s:\n", err)
		}

		// 1. If you comment first way, and then type err here, the go garbage collector discards it, it does not recognise it.
		// 2. But if you comment second way, and then type err, it is defined and err acts as a local variable.
	}
}
