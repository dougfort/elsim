package main

type elevatorStruct struct {
	currentFloor int
	curretState  ElevatorState

	// the elevator can have multiple stops requested
	destFloors []int

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

// StartTrip starts a new trip for the elevator
// It opens the doors to receive occupants
func (e *elevatorStruct) StartTrip(toFloors []int) {
	e.destFloors = toFloors // these should be in order
	e.curretState = ElevatorDoorsOpen
}

// AddFloor adds a floor to an existing trip
func (e *elevatorStruct) AddFloor(toFloor int)

// PerformMaintenence brings the elevator from ElevatorWaitingMaintenence
// to ElevatorIdle at the current floor
func (e *elevatorStruct) PerformMaintenence() {
	e.tripCount = 0
	e.curretState = ElevatorIdle
}
