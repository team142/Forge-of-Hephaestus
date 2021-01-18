package repository

func NewState() *State {
	result := &State{
		HighwayRepo: NewHighwayRepository(),
		OffloadRepo: NewOffloadRepository(),
		ParkingRepo: NewParkingRepository(),
		RefuelRepo:  NewRefuelRepository(),
		TurtleRepo:  NewTurtleRepository(),
	}
	return result
}

type State struct {
	HighwayRepo *GenericNamedLocationRepository
	OffloadRepo *GenericNamedLocationRepository
	ParkingRepo *GenericNamedLocationRepository
	RefuelRepo  *GenericNamedLocationRepository
	TurtleRepo  *TurtleRepository
}
