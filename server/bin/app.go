package main

import (
	"github.com/sirupsen/logrus"
	"github.com/team142/Forge-of-Hephaestus/server/controller"
	"github.com/team142/Forge-of-Hephaestus/server/repository"
	"github.com/team142/Forge-of-Hephaestus/server/web"
)

func main() {
	logrus.Infoln("Starting up!")
	logrus.Println("> Loading repositories")
	ServerState := repository.NewState()

	controller.Orchestrate(ServerState)
	web.StartServer(":8080")

}
