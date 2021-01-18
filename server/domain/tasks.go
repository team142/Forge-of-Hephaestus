package domain

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
	Done              bool
}

type TaskType string

const (
	TaskPark             TaskType = "TASK_PARK"
	TaskLeaveGarage      TaskType = "TASK_LEAVE_GARAGE"
	TaskTravelToX        TaskType = "TASK_TRAVEL_TO_X"
	TaskStartKelpCut     TaskType = "TASK_START_KELP_CUT"
	TaskStartKelpCollect TaskType = "TASK_START_KELP_COLLECT"
	TaskDropOffItems     TaskType = "TASK_DROP_OFF_ITEMS"
	TaskRefuel           TaskType = "TASK_REFUEL"
)

func IsTaskRecoverable(taskType TaskType) bool {
	found, recoverable := Recoverability[taskType]
	if !found {
		return false
	}
	return recoverable
}

var Recoverability = map[TaskType]bool{
	TaskPark:             true,
	TaskLeaveGarage:      true,
	TaskTravelToX:        false,
	TaskStartKelpCut:     false,
	TaskStartKelpCollect: false,
	TaskDropOffItems:     true,
	TaskRefuel:           true,
}
