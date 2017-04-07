package main

type controllerStruct struct {
	elevators []Elevator
}

// NewController returns an entity tat implements the Controller interface
func NewController(elevators []Elevator) Controller {
	return &controllerStruct{
		elevators: elevators,
	}
}

// Step takes a single step
func (c *controllerStruct) Step(stepCount int) {

}
