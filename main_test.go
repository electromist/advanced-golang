package main

import (
	"testing"
)

func TestMain(t *testing.T) {
	t.Run("processTruck", func(t *testing.T) {
		t.Run("should load and unload a truck cargo", func(t *testing.T) {
			// Initialize
			nt := &NormalTruck{id: "1", cargo: 42}
			et := &ElectricTruck{id: "2"}

			// Process
			if err := processTruck(nt); err != nil {
				t.Fatalf("Error processing normal truck: %s", err)
			}
			if err := processTruck(et); err != nil {
				t.Fatalf("Error processing electric truck: %s", err)
			}

			// Assertions
			if nt.cargo != 0 {
				t.Fatal("Normal truck cargo should be 0")
			}

			if et.battery != -2 {
				t.Fatal("Electric truck battery should be -2")
			}
		})
	})
}
