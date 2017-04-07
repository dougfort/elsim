package main

type controllerStruct struct {
}

// NewController returns an entity tat implements the Controller interface
func NewController() Controller {
	return &controllerStruct{}
}

// Step takes a single step
func (c *controllerStruct) Step(stepCount int) {

}
