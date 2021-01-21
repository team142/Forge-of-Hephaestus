package domain

import "github.com/google/uuid"

/*
	Tasks are a list of things that happen to meet a job
*/

type Task struct {
	ID                string
	TaskType          TaskType
	ActualWork        bool
	Recoverable       bool //Ensures we skip unrecoverable tasks after server reset or chunk reload
	MetaIntegers      map[string]int
	Meta              map[string]string
	RealizedAsActions bool
	TaskContext       TaskContext
	Actions           []Action
	Done              bool
}

type TaskType string

const (
	TaskPark             TaskType = "TASK_PARK"
	TaskLeaveGarage      TaskType = "TASK_LEAVE_GARAGE"
	TaskCollectSeeds     TaskType = "TASK_COLLECT_SEEDS"
	TaskTravelToX        TaskType = "TASK_TRAVEL_TO_X"
	TaskStartWheatCut    TaskType = "TASK_START_WHEAT_CUT"
	TaskStartKelpCut     TaskType = "TASK_START_KELP_CUT"
	TaskStartKelpCollect TaskType = "TASK_START_KELP_COLLECT"
	TaskDropOffItems     TaskType = "TASK_DROP_OFF_ITEMS"
	TaskRefuel           TaskType = "TASK_REFUEL"
	TaskTurnToDirection  TaskType = "TASK_TURN_TO_DIRECTION"
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
	TaskCollectSeeds:     false,
	TaskStartWheatCut:    false,
	TaskTurnToDirection:  true,
	TaskStartKelpCut:     false,
	TaskStartKelpCollect: false,
	TaskDropOffItems:     true,
	TaskRefuel:           true,
}

var TaskPlanners = map[TaskType]func(orchestration Orchestration, ctx TaskContext) Task{
	TaskPark:             NewTaskPark,
	TaskLeaveGarage:      NewTaskLeaveGarage,
	TaskTravelToX:        NewTaskTravelToX,
	TaskCollectSeeds:     NewTaskCollectSeeds,
	TaskStartWheatCut:    NewTaskWheatCut,
	TaskTurnToDirection:  NewTaskTurnToDirection,
	TaskStartKelpCut:     NewTaskStartKelpCut,
	TaskStartKelpCollect: NewTaskStartKelpCollect,
	TaskDropOffItems:     NewTaskDropOffItems,
	TaskRefuel:           NewTaskRefuel,
}

type TaskContext struct {
	TypedNamedLocation TypedNamedLocation
	Direction          Direction
}

func NewTask(taskType TaskType, ctx TaskContext) Task {
	return Task{
		ID:                uuid.New().String(),
		TaskType:          taskType,
		ActualWork:        false,
		Recoverable:       false,
		TaskContext:       ctx,
		MetaIntegers:      map[string]int{},
		Meta:              map[string]string{},
		RealizedAsActions: false,
		Actions:           make([]Action, 0),
		Done:              false,
	}
}

func NewTaskPark(orchestration Orchestration, ctx TaskContext) Task {
	return NewTask(TaskPark, ctx)
}

func NewTaskTurnToDirection(orchestration Orchestration, ctx TaskContext) Task {
	return NewTask(TaskTurnToDirection, ctx)
}

func NewTaskLeaveGarage(orchestration Orchestration, ctx TaskContext) Task {
	result := NewTask(TaskLeaveGarage, ctx)
	return result
}

func NewTaskTravelToX(orchestration Orchestration, ctx TaskContext) Task {
	result := NewTask(TaskTravelToX, ctx)
	result.Recoverable = true
	result.MetaIntegers = orchestration.Spec.MetaIntegers
	return result
}

func NewTaskWheatCut(orchestration Orchestration, ctx TaskContext) Task {
	result := NewTask(TaskStartWheatCut, ctx)
	return result

}

func NewTaskCollectSeeds(orchestration Orchestration, ctx TaskContext) Task {
	result := NewTask(TaskCollectSeeds, ctx)
	return result
}

func NewTaskStartKelpCut(orchestration Orchestration, ctx TaskContext) Task {
	return NewTask(TaskStartKelpCut, ctx)
}

func NewTaskStartKelpCollect(orchestration Orchestration, ctx TaskContext) Task {
	return NewTask(TaskStartKelpCollect, ctx)
}

func NewTaskDropOffItems(orchestration Orchestration, ctx TaskContext) Task {
	return NewTask(TaskDropOffItems, ctx)
}

func NewTaskRefuel(orchestration Orchestration, ctx TaskContext) Task {
	return NewTask(TaskRefuel, ctx)
}
