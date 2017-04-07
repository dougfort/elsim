package main

// TODO: get control constants from commandline and/or config file
const steps = 1
const elevators = 1
const floors = 1

// Stepper implements single steps in the simulation
type Stepper interface {
	Step(stepCount int)
}

// Controller controls the elevators
type Controller interface {
	Stepper
}

// Elevator simulates a single elevator
type Elevator interface {
	Steper
}
