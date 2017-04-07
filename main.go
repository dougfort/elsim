package main

func main() {
	// TODO: get simData from a file or some eternal souce
	// simData drives the simualation. Each entry in the outer slice represents
	// one step in the simulation
	simData := [][]FloorRequest{
		[]FloorRequest{FloorRequest{1, 2}},
		[]FloorRequest{},
	}

	var elevators []Elevator
	for i := 0; i < TotalElevators; i++ {
		// we could start with elevators randomly distributed by flor and state
		elevators = append(elevators, NewElevator(i, 1, ElevatorIdle))
	}
	controller := NewController(elevators)

	// the simulation runs in discrete 'steps'
	// we give the controller some floor requests (possibly 0)
	// and then take a step
	for step, stepData := range simData {
		for _, request := range stepData {
			controller.AcceptReqest(request)
		}
		controller.Step(step)
		for _, elevator := range elevators {
			elevator.Step(step)
		}
	}
}
