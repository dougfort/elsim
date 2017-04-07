package main

func main() {
	// TODO: get simData from a file or some eternal souce
	// simData drives the simualation. Each entry in the outer slice represents
	// one step in the simulation
	simData := [][]FloorRequest{
		[]FloorRequest{FloorRequest{1, 2}},
		[]FloorRequest{},
	}

	elevators := make([]Elevator, totalElevators)
	controller := NewController(elevators)

	// the simulation runs in discrete 'steps'
	// we tell each object to step
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
