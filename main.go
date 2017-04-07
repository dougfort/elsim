package main

func main() {

	elevators := make([]Elevator, totalElevators)
	controller := NewController(elevators)

	// the simulation runs in discrete 'steps'
	// we tell each object to step
	for step := 0; step < totalSteps; step++ {
		controller.Step(step)
		for _, elevator := range elevators {
			elevator.Step(step)
		}
	}
}
