package repository

import (
	"github.com/team142/Forge-of-Hephaestus/server/domain"
	"sync"
)

func NewTurtleRepository() *TurtleRepository {
	result := &TurtleRepository{
		turtles: make(map[string]*domain.Turtle),
	}
	return result
}

type TurtleRepository struct {
	sync.Mutex
	turtles map[string]*domain.Turtle
}

func (t *TurtleRepository) OverwriteTurtle(newTurtle *domain.Turtle) {
	t.Lock()
	defer t.Unlock()
	t.turtles[newTurtle.ID] = newTurtle
}

func (t *TurtleRepository) ForEachTurtle(f func(turtle *domain.Turtle)) {
	t.Lock()
	defer t.Unlock()
	for _, next := range t.turtles {
		f(next)
	}
}

func (t *TurtleRepository) GetAllTurtles() map[string]*domain.Turtle {
	t.Lock()
	defer t.Unlock()
	result := t.turtles
	return result
}
