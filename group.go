package main

// Group represents a group of people.
// Group can be traveling or waiting for a car.
type Group struct {
	id     int
	people int
	car    *Car
}

// isTraveling checks if Group is traveling.
func (g *Group) isTraveling() bool {
	return g.car != nil
}

// isWaiting checks if Group is waiting for a car.
func (g *Group) isWaiting() bool {
	return g.car == nil
}

// endTravel removes associated car to the group ending the travel.
func (g *Group) endTravel() {
	if g.car != nil {
		g.car.group = nil
	}

	g.car = nil
}

// startTravel starts travel in given Car.
func (g *Group) startTravel(c *Car) {
	g.car = c
	g.car.group = g
}
