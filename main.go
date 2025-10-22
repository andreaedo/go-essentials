package main

import (
	"errors"
	"fmt"
	"log"
)

var (
	ErrNotImplemented = errors.New("not implemented")
	ErrTruckNotfound  = errors.New("truck not found")
)

type Truck struct {
	id string
}

func (t *Truck) LoadCargo() error {
	return nil
}

func (t *Truck) UnloadCargo() error {
	return nil

}

func processTruck(truck Truck) error {
	fmt.Printf("Processing truck: %s\n", truck.id)

	if err := truck.LoadCargo(); err != nil {
		return fmt.Errorf("error Loading Cargo: %w", err)
	}

	if err := truck.UnloadCargo(); err != nil {
		return fmt.Errorf("error Unloading Cargo: %w", err)

	}

	return nil
}

func main() {
	trucks := []Truck{
		{id: "Truck-1"},
		{id: "Truck-2"},
		{id: "Truck-3"},
	}

	for _, truck := range trucks {
		fmt.Printf("Truck %s arrived.\n", truck.id)

		if err := processTruck(truck); err != nil {
			log.Fatalf("Error processing truck: %s\n", err)
			//switch err {
			//case ErrNotImplemented:
			//	return
			//case ErrTruckNotfound:
			//	return
			//default:
			//	fmt.Printf("Error processing truck: %s\n", err)
			//}
		}
	}
}
