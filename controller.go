package main

type controllerStruct struct {
	elevators       []Elevator
	pendingRequests []FloorRequest
	activeRequests  map[int]FloorRequest
}

// NewController returns an entity tat implements the Controller interface
func NewController(elevators []Elevator) Controller {
	return &controllerStruct{
		elevators:      elevators,
		activeRequests: make(map[int]FloorRequest),
	}
}

// Step takes a single step
func (c *controllerStruct) Step(stepCount int) {

	// check the elevator state machines
	for i, elevator := range c.elevators {
		switch elevator.CurrentState() {
		case ElevatorIdle:
			// if the elevator became idle after an activeRequest,
			// it has completed the request
			delete(c.activeRequests, i)
		case ElevatorWaitingMaintenence:
			elevator.PerformMaintenence()
		}
	}

	// see if we can start any new requests
	var newPending []FloorRequest
PENDING_LOOP:
	for _, request := range c.pendingRequests {

		// if there is an unoccupied elevator on the requesting floor, use it
		for i, elevator := range c.elevators {
			if elevator.CurrentFloor() == request.From &&
				elevator.CurrentState() == ElevatorIdle {
				elevator.StartTrip(request.To)
				c.activeRequests[i] = request
			}
			continue PENDING_LOOP
		}

		// if an occupied elevator is moving and will pass the From floor on
		// its way, use that one
	}

}

// Accept a request to go from one floor to another
func (c *controllerStruct) AcceptReqest(request FloorRequest) {
	c.pendingRequests = append(c.pendingRequests, request)
}
