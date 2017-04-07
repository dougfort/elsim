package main

// TODO: get control constants from commandline and/or config file
const totalSteps = 1
const totalElevators = 1
const totalFloors = 1

// Stepper implements single steps in the simulation
type Stepper interface {
	Step(stepCount int)
}

// FloorRequest
type FloorRequest struct {
}

// Controller controls the elevators
type Controller interface {
	Stepper
}

// Elevator simulates a single elevator
type Elevator interface {
	Stepper

	// CurrentFloor reports the elevator's location
	CurrentFloor() int
}
