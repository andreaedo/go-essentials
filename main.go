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

type Truck interface {
	LoadCargo() error
	UnloadCargo() error
}

type NormalTruck struct {
	id    string
	cargo int
}

func (t *NormalTruck) LoadCargo() error {
	t.cargo += 1
	return nil
}

func (t *NormalTruck) UnloadCargo() error {
	t.cargo = 0
	return nil
}

type ElectricTruck struct {
	id      string
	cargo   int
	battery float64
}

func (e *ElectricTruck) LoadCargo() error {
	e.cargo += 1
	e.battery += -1
	return nil
}

func (e *ElectricTruck) UnloadCargo() error {
	e.cargo = 0
	e.battery += -1
	return nil
}

func processTruck(truck Truck) error {
	fmt.Printf("processing truck %+v\n", truck)

	if err := truck.LoadCargo(); err != nil {
		return fmt.Errorf("error Loading Cargo: %w", err)
	}

	if err := truck.UnloadCargo(); err != nil {
		return fmt.Errorf("error Unloading Cargo: %w", err)
	}

	return nil
}

func main() {
	// testing git credentials
	nt := &NormalTruck{id: "1"}
	et := &ElectricTruck{id: "2"}

	if err := processTruck(nt); err != nil {
		log.Fatalf("Error processing truck: %s\n", err)
	}

	if err := processTruck(et); err != nil {
		log.Fatalf("Error processing truck: %s\n", err)
	}

	log.Println(nt.cargo)
	log.Println(et.battery)

}
