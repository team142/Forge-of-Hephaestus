package repository

import (
	"github.com/team142/Forge-of-Hephaestus/server/domain"
	"github.com/team142/Forge-of-Hephaestus/server/persistence"
)

const orchestrationFilesPath = "state/orch"

func NewState() *State {
	result := &State{
		NamedRepositories: map[StateType]*GenericNamedLocationRepository{
			StateTypeHighwayRepo: NewHighwayRepository(),
			StateTypeOffloadRepo: NewOffloadRepository(),
			StateTypeParkingRepo: NewParkingRepository(),
			StateTypeRefuelRepo:  NewRefuelRepository(),
		},
		TurtleRepo: NewTurtleRepository(),
		OrchestrationRepository: NewOrchestrationRepository(
			orchestrationFilesPath,
			persistence.RepositoryOrchestrationsInitFromDisk,
			persistence.RepositoryOrchestrationsPersistToDisk,
		),
	}
	domain.GlobalLocations = domain.StaticLocations{
		TNLDropOff:         result.NamedRepositories[StateTypeOffloadRepo].GetAll()[0],
		TNLRefuel:          result.NamedRepositories[StateTypeRefuelRepo].GetAll()[0],
		TNLLeaveGarageSpot: domain.NamedLocation{},
		TNLGarageEntrance:  domain.NamedLocation{},
	}
	return result
}

type State struct {
	NamedRepositories       map[StateType]*GenericNamedLocationRepository
	TurtleRepo              *TurtleRepository
	OrchestrationRepository *OrchestrationRepository
}

func (s *State) GetHighway() *GenericNamedLocationRepository {
	return s.NamedRepositories[StateTypeHighwayRepo]
}
func (s *State) GetOffload() *GenericNamedLocationRepository {
	return s.NamedRepositories[StateTypeOffloadRepo]
}
func (s *State) GetParking() *GenericNamedLocationRepository {
	return s.NamedRepositories[StateTypeParkingRepo]
}
func (s *State) GetRefuel() *GenericNamedLocationRepository {
	return s.NamedRepositories[StateTypeRefuelRepo]
}

type StateType string

const (
	StateTypeHighwayRepo StateType = "Highway"
	StateTypeOffloadRepo StateType = "Offload"
	StateTypeParkingRepo StateType = "Parking"
	StateTypeRefuelRepo  StateType = "Refuel"
)
