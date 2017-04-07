package main

// TODO: get control constants from commandline and/or config file
const (
	TotalElevators             = 1
	TotalFloors                = 1
	MaxTripsBetweenMaintenence = 100
)

// ElevatorState identifies the current state of an elevator
type ElevatorState int

const (
	ElevatorIdle ElevatorState = iota
	ElevatorWaitingMaintenence
	ElevatorDoorsOpen
	ElevatorDoorsClosed
	ElevatorMovingUp
	ElevatorMovingDown
)

// Stepper implements single steps in the simulation
type Stepper interface {
	Step(stepCount int)
}

// FloorRequest is a request to got from one floor to another
type FloorRequest struct {
	From int
	To   int
}

// Controller controls the elevators
type Controller interface {
	Stepper

	// Accept a request to go from one floor to another
	AcceptReqest(request FloorRequest)
}

// Elevator simulates a single elevator
type Elevator interface {
	Stepper

	// CurrentFloor reports the elevator's location
	CurrentFloor() int

	// CurrentState reports the Elevator's current state
	CurrentState() ElevatorState

	// PerformMaintenence brings the elevator from ElevatorWaitingMaintenence
	// to ElevatorIdle at the current floor
	PerformMaintenence()
}
