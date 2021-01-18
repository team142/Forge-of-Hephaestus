package controller

import (
	"github.com/team142/Forge-of-Hephaestus/server/repository"
)

func Orchestrate(state *repository.State) {
	startScheduler(state)
}
