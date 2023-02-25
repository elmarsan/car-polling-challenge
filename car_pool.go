package main

import (
	"log"
	"sync"
)

// CarPool represents all cars of the system.
type CarPool struct {
	cars    []*Car
	waiting []*Group

	releasedCarChan chan *Car
	mutex           sync.Mutex
}

// NewCarPool returns CarPool with given cars.
func NewCarPool(cars []*Car) *CarPool {
	return &CarPool{
		cars:            cars,
		waiting:         []*Group{},
		releasedCarChan: make(chan *Car),
		mutex:           sync.Mutex{},
	}
}

func (p *CarPool) start() {
	for {
		log.Printf("Waiting for a free car...\n")

		c := <-p.releasedCarChan

		log.Printf("Car %d it's available\n", c.id)

		p.mutex.Lock()

		for _, g := range p.waiting {
			if c.seats >= g.people {
				log.Printf("Assign car %d to group %d\n", c.id, g.id)
				c.startTravel(g)
			}
		}

		p.mutex.Unlock()
	}
}

func (p *CarPool) addGroup(g *Group) {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	c := p.findFree(g.people)
	if c != nil {
		log.Printf("Group %d started traveling on car %d", g.id, c.id)
		c.startTravel(g)
	} else {
		log.Printf("Group %d added to waiting queue", g.id)
		p.waiting = append(p.waiting, g)
	}
}

// findFree finds waiting car with car.seats >= seats.
func (p *CarPool) findFree(seats int) *Car {
	for _, c := range p.cars {
		if c.seats >= seats && c.isWaiting() {
			return c
		}
	}

	return nil
}

func (p *CarPool) dropGroup(id int) bool {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	var releasedCar *Car = nil

	for _, c := range p.cars {
		if c.group != nil && c.group.id == id {
			releasedCar = c
			c.group = nil
			log.Printf("Group %d dropped succesfully", id)
		}
	}

	if releasedCar != nil {
		p.releasedCarChan <- releasedCar
		return true
	}

	index := -1
	for i, g := range p.waiting {
		if g.id == id {
			index = i
			break
		}
	}

	if index != -1 {
		p.waiting = append(p.waiting[:index], p.waiting[index+1:]...)
		return true
	}

	log.Printf("Group %d not found", id)
	return false
}
