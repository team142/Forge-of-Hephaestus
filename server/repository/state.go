package repository

import "github.com/team142/Forge-of-Hephaestus/server/persistence"

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
