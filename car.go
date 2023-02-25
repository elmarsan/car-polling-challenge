package main

// Car represents a car.
// Car can being traveling carrying a group or waiting.
type Car struct {
	id    int
	seats int
	group *Group
}

// isTraveling checks if Car is traveling carrying a group.
func (c *Car) isTraveling() bool {
	return c.group != nil
}

// isWaiting checks if Car is waiting for travel.
func (c *Car) isWaiting() bool {
	return c.group == nil
}

// endTravel removes associated group to the car ending the travel
func (c *Car) endTravel() {
	if c.group != nil {
		c.group.car = nil
	}

	c.group = nil
}

// startTravel starts travel with given Group.
func (c *Car) startTravel(g *Group) {
	c.group = g
	c.group.car = c
}
