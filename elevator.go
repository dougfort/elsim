package main

type elevatorStruct struct {
	currentFloor int
	curretState  ElevatorState
	destFloor    int

	// a trip is counted when the elevator opens its doors on a destination floor
	tripCount int
}

// NewElevator returns an entity tat implements the Elevator interface
func NewElevator(
	startingFloor int,
	startingState ElevatorState,
) Elevator {
	return &elevatorStruct{
		currentFloor: startingFloor,
		curretState:  startingState,
	}
}

// Step takes a single step
func (e *elevatorStruct) Step(stepCount int) {
}

// CurrentFloor reports the elevator's location
func (e *elevatorStruct) CurrentFloor() int {
	return e.currentFloor
}

// CurrentState reports the Elevator's current state
func (e *elevatorStruct) CurrentState() ElevatorState {
	return e.curretState
}

// PerformMaintenence brings the elevator from ElevatorWaitingMaintenence
// to ElevatorIdle at the current floor
func (e *elevatorStruct) PerformMaintenence() {
	e.tripCount = 0
	e.curretState = ElevatorIdle
}
