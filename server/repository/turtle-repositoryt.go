package repository

import (
	"github.com/team142/Forge-of-Hephaestus/server/model"
	"sync"
)

func NewTurtleRepository() *TurtleRepository {
	result := &TurtleRepository{
		turtles: make(map[string]*model.Turtle),
	}
	return result
}

type TurtleRepository struct {
	sync.Mutex
	turtles map[string]*model.Turtle
}

func (t *TurtleRepository) OverwriteTurtle(newTurtle *model.Turtle) {
	t.Lock()
	defer t.Unlock()
	t.turtles[newTurtle.ID] = newTurtle
}

func (t *TurtleRepository) ForEachTurtle(f func(turtle *model.Turtle)) {
	t.Lock()
	defer t.Unlock()
	for _, next := range t.turtles {
		f(next)
	}
}

func (t *TurtleRepository) GetAllTurtles() map[string]*model.Turtle {
	t.Lock()
	defer t.Unlock()
	result := t.turtles
	return result
}
