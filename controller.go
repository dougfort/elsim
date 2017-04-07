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

}

// Accept a request to go from one floor to another
func (c *controllerStruct) AcceptReqest(request FloorRequest) {
	c.pendingRequests = append(c.pendingRequests, request)
}
