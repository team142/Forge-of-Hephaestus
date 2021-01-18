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
	HighwayRepo *GenericRepository
	OffloadRepo *GenericRepository
	ParkingRepo *GenericRepository
	RefuelRepo  *GenericRepository
	TurtleRepo  *TurtleRepository
}
