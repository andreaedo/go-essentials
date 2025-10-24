package main

import (
	"testing"
)

func TestTruck(t *testing.T) {
	t.Run("processTruck", func(t *testing.T) {
		t.Run("should load and unload a truck cargo", func(t *testing.T) {
			nt := &NormalTruck{id: "1", cargo: 42}
			et := &ElectricTruck{id: "2"}

			if err := processTruck(nt); err != nil {
				t.Fatalf("Error processing truck: %s\n", err)
			}

			if err := processTruck(et); err != nil {
				t.Fatalf("Error processing truck: %s\n", err)
			}

			//asserting
			if nt.cargo != 0 {
				t.Fatal("normal truck cargo should be 0")
			}

			if et.battery != -2 {
				t.Fatal("electric truck battery should be -2")
			}

		})
	})
}
