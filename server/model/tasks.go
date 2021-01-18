package model

/*
	Tasks are a list of things that happen to meet a job
*/

type Task struct {
	ID                string
	TaskType          TaskType
	ActualWork        bool
	Recoverable       bool //Ensures we skip unrecoverable tasks after server reset or chunk reload
	PreStartLocation  NamedLocation
	RealizedAsActions bool
	Actions           []Action
}

type TaskType string

const (
	TaskReturnGarage     TaskType = "TASK_RETURN_GARAGE"
	TaskLeaveGarage      TaskType = "TASK_LEAVE_GARAGE"
	TaskGetOnHighway     TaskType = "TASK_GET_ON_HIGHWAY"
	TaskTravelOnHighway  TaskType = "TASK_TRAVEL_ON_HIGHWAY"
	TaskStartKelpCut     TaskType = "TASK_START_KELP_CUT"
	TaskStartKelpCollect TaskType = "TASK_START_KELP_COLLECT"
	TaskDropOffItems     TaskType = "TASK_DROP_OFF_ITEMS"
	TaskRefuel           TaskType = "TASK_REFUEL"
	TaskPark             TaskType = "TASK_PARK"
)

func IsTaskRecoverable(taskType TaskType) bool {
	found, recoverable := Recoverability[taskType]
	if !found {
		return false
	}
	return recoverable
}

var Recoverability = map[TaskType]bool{
	TaskReturnGarage:     true,
	TaskLeaveGarage:      true,
	TaskGetOnHighway:     true,
	TaskTravelOnHighway:  false,
	TaskStartKelpCut:     false,
	TaskStartKelpCollect: false,
	TaskDropOffItems:     true,
	TaskRefuel:           true,
	TaskPark:             true,
}
