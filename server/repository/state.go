package repository

import (
	"github.com/team142/Forge-of-Hephaestus/server/domain"
	"github.com/team142/Forge-of-Hephaestus/server/persistence"
)

const orchestrationFilesPath = "state/orch"

func NewState() *State {
	result := &State{
		HighwayRepo: NewHighwayRepository(),
		OffloadRepo: NewOffloadRepository(),
		ParkingRepo: NewParkingRepository(),
		RefuelRepo:  NewRefuelRepository(),
		TurtleRepo:  NewTurtleRepository(),
		OrchestrationRepository: NewOrchestrationRepository(
			orchestrationFilesPath,
			persistence.RepositoryOrchestrationsInitFromDisk,
			persistence.RepositoryOrchestrationsPersistToDisk,
		),
	}
	domain.GlobalLocations = domain.StaticLocations{
		TNLDropOff:         result.OffloadRepo.GetAll()[0],
		TNLRefuel:          result.RefuelRepo.GetAll()[0],
		TNLLeaveGarageSpot: domain.NamedLocation{},
		TNLGarageEntrance:  domain.NamedLocation{},
	}
	return result
}

type State struct {
	HighwayRepo             *GenericNamedLocationRepository
	OffloadRepo             *GenericNamedLocationRepository
	ParkingRepo             *GenericNamedLocationRepository
	RefuelRepo              *GenericNamedLocationRepository
	TurtleRepo              *TurtleRepository
	OrchestrationRepository *OrchestrationRepository
}
