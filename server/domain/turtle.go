package domain

import (
	"sync"
	"time"
)

type Turtle struct {
	sync.Mutex
	ID           string
	features     []TurtleFeature
	x, y, z      int
	direction    Direction
	waiting      bool
	lastCheckIn  time.Time
	job          *Job
	turtleStatus TurtleStatus
}

type TurtleStatus string

const (
	TurtleStatusNull          TurtleStatus = "TURTLE_STATUS_LOST"
	TurtleStatusExecutingTask TurtleStatus = "TURTLE_STATUS_EXECUTING_TASK"
	TurtleStatusAvailable     TurtleStatus = "TURTLE_STATUS_AVAILABLE"
)

func (t Turtle) LastCheckIn() time.Time {
	t.Lock()
	defer t.Unlock()
	return t.lastCheckIn
}

func (t Turtle) Waiting() bool {
	t.Lock()
	defer t.Unlock()
	return t.waiting
}

func (t Turtle) Direction() Direction {
	t.Lock()
	defer t.Unlock()
	return t.direction
}

func (t Turtle) Z() int {
	t.Lock()
	defer t.Unlock()
	return t.z
}
func (t Turtle) X() int {
	t.Lock()
	defer t.Unlock()
	return t.x
}
func (t Turtle) Y() int {
	t.Lock()
	defer t.Unlock()
	return t.y
}

func (t Turtle) Features() []TurtleFeature {
	t.Lock()
	defer t.Unlock()
	return t.features
}

func (t Turtle) HasFeature(feature TurtleFeature) bool {
	t.Lock()
	defer t.Unlock()
	for _, next := range t.features {
		if next == feature {
			return true
		}
	}
	return false
}
