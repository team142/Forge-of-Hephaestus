package controller

import (
	"github.com/team142/Forge-of-Hephaestus/server/repository"
	"time"
)

func startScheduler(state *repository.State) {
	go func() {
		for {
			//TODO: check if jobs are outstanding

			time.Sleep(1 * time.Second)
		}
	}()
}
