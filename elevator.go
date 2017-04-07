package main

type elevatorStruct struct {
	elevatorID   int
	currentFloor int
}

// NewElevator returns an entity tat implements the Elevator interface
func NewElevator(elevatorID int, startingFloor int) Elevator {
	return &elevatorStruct{
		elevatorID:   elevatorID,
		currentFloor: startingFloor,
	}
}

// Step takes a single step
func (e *elevatorStruct) Step(stepCount int) {

}

// CurrentFloor reports the elevator's location
func (e *elevatorStruct) CurrentFloor() int {
	return e.currentFloor
}
