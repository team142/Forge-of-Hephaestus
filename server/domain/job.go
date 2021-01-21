package domain

import (
	"fmt"
)

/*
	Overarching job is a particular service carried out by Turtles
	A consists of a list of tasks. These are determined by its dynamic target
*/
type Job struct {
	Name              JobType
	Tasks             []Task
	CurrentTaskID     string
	OrchestrationName string
}

func (j Job) IsActualWorkDone() bool {
	for _, t := range j.Tasks {
		if t.ActualWork && t.Done {
			return true
		}
	}
	return false
}

type JobType string

const (
	JobCutKelp        JobType = "JOB_CUT_KELP"
	JobCollectKelp    JobType = "JOB_COLLECT_KELP"
	JobHarvestWheat   JobType = "JOB_HARVEST_WHEAT"
	JobSortChests     JobType = "JOB_SORT_CHESTS"
	JobCraftBread     JobType = "JOB_CRAFT_BREAD"
	JobCraftKelpBlock JobType = "JOB_CRAFT_KELP_BLOCK"
)

// JobPlanners accept a *Job and populate the tasks
var JobPlanners = map[JobType]func(*Turtle, *Job, Orchestration){
	JobHarvestWheat: planHarvestWheat,
	JobCutKelp:      planCutKelp,
	JobCollectKelp:  planCollectKelp,
}

func NewJob(turtle *Turtle, orchestration Orchestration) *Job {
	result := &Job{
		Name:              orchestration.JobType,
		Tasks:             make([]Task, 0),
		CurrentTaskID:     "",
		OrchestrationName: orchestration.Name,
	}
	populateTypedNamedLocations(turtle, &orchestration)
	handler, found := JobPlanners[orchestration.JobType]
	if !found {
		panic(fmt.Sprint("no job planner found for ", orchestration.JobType))
	}
	handler(turtle, result, orchestration)
	return result
}

func populateTypedNamedLocations(turtle *Turtle, orchestration *Orchestration) {

	orchestration.Spec.TypedNamedLocations[TNLLeaveGarageSpot] = GlobalLocations.TNLLeaveGarageSpot
	orchestration.Spec.TypedNamedLocations[TNLDropOff] = GlobalLocations.TNLDropOff
	orchestration.Spec.TypedNamedLocations[TNLRefuel] = GlobalLocations.TNLRefuel
	orchestration.Spec.TypedNamedLocations[TNLGarageEntrance] = GlobalLocations.TNLGarageEntrance

	//Assigned to that Bot
	orchestration.Spec.TypedNamedLocations[TNLParking] = turtle.parking

}

func planHarvestWheat(t *Turtle, j *Job, orchestration Orchestration) {

	//Leave Garage
	j.Tasks = append(j.Tasks, TaskPlanners[TaskLeaveGarage](orchestration, TaskContext{TypedNamedLocation: TNLLeaveGarageSpot}))

	//Seeds
	j.Tasks = append(j.Tasks, TaskPlanners[TaskTravelToX](orchestration, TaskContext{TypedNamedLocation: TNLSeeds}))
	j.Tasks = append(j.Tasks, TaskPlanners[TaskCollectSeeds](orchestration, TaskContext{}))

	//Wheat
	j.Tasks = append(j.Tasks, TaskPlanners[TaskTravelToX](orchestration, TaskContext{TypedNamedLocation: TNLWheat}))
	j.Tasks = append(j.Tasks, TaskPlanners[TaskStartWheatCut](orchestration, TaskContext{}))

	//Drop off
	j.Tasks = append(j.Tasks, TaskPlanners[TaskTravelToX](orchestration, TaskContext{TypedNamedLocation: TNLDropOff}))
	j.Tasks = append(j.Tasks, TaskPlanners[TaskDropOffItems](orchestration, TaskContext{}))

	//Refuel
	j.Tasks = append(j.Tasks, TaskPlanners[TaskTravelToX](orchestration, TaskContext{TypedNamedLocation: TNLRefuel}))
	j.Tasks = append(j.Tasks, TaskPlanners[TaskRefuel](orchestration, TaskContext{}))

	//Go home
	j.Tasks = append(j.Tasks, TaskPlanners[TaskTravelToX](orchestration, TaskContext{TypedNamedLocation: TNLGarageEntrance}))
	j.Tasks = append(j.Tasks, TaskPlanners[TaskPark](orchestration, TaskContext{TypedNamedLocation: TNLParking}))
	j.Tasks = append(j.Tasks, TaskPlanners[TaskTurnToDirection](orchestration, TaskContext{Direction: t.parkingDirection}))

}

func planCutKelp(turtle *Turtle, job *Job, orchestration Orchestration) {
	//Leave Garage
	job.Tasks = append(job.Tasks, TaskPlanners[TaskLeaveGarage](orchestration, TaskContext{TypedNamedLocation: TNLLeaveGarageSpot}))

	//Cut kelp
	job.Tasks = append(job.Tasks, TaskPlanners[TaskTravelToX](orchestration, TaskContext{TypedNamedLocation: TNLCutKelp}))
	job.Tasks = append(job.Tasks, TaskPlanners[TaskStartKelpCut](orchestration, TaskContext{}))

	//Drop off
	job.Tasks = append(job.Tasks, TaskPlanners[TaskTravelToX](orchestration, TaskContext{TypedNamedLocation: TNLDropOff}))
	job.Tasks = append(job.Tasks, TaskPlanners[TaskDropOffItems](orchestration, TaskContext{})) //TODO: becomes DropOffFurnace

	//Refuel
	job.Tasks = append(job.Tasks, TaskPlanners[TaskTravelToX](orchestration, TaskContext{TypedNamedLocation: TNLRefuel}))
	job.Tasks = append(job.Tasks, TaskPlanners[TaskRefuel](orchestration, TaskContext{}))

	//Go home
	job.Tasks = append(job.Tasks, TaskPlanners[TaskTravelToX](orchestration, TaskContext{TypedNamedLocation: TNLGarageEntrance}))
	job.Tasks = append(job.Tasks, TaskPlanners[TaskPark](orchestration, TaskContext{TypedNamedLocation: TNLParking}))
	job.Tasks = append(job.Tasks, TaskPlanners[TaskTurnToDirection](orchestration, TaskContext{Direction: turtle.parkingDirection}))

}

func planCollectKelp(turtle *Turtle, job *Job, orchestration Orchestration) {
	//Leave Garage
	job.Tasks = append(job.Tasks, TaskPlanners[TaskLeaveGarage](orchestration, TaskContext{TypedNamedLocation: TNLLeaveGarageSpot}))

	//Cut kelp
	job.Tasks = append(job.Tasks, TaskPlanners[TaskTravelToX](orchestration, TaskContext{TypedNamedLocation: TNLCollectKelp}))
	job.Tasks = append(job.Tasks, TaskPlanners[TaskStartKelpCollect](orchestration, TaskContext{}))

	//Drop off
	job.Tasks = append(job.Tasks, TaskPlanners[TaskTravelToX](orchestration, TaskContext{TypedNamedLocation: TNLDropOff}))
	job.Tasks = append(job.Tasks, TaskPlanners[TaskDropOffItems](orchestration, TaskContext{}))

	//Refuel
	job.Tasks = append(job.Tasks, TaskPlanners[TaskTravelToX](orchestration, TaskContext{TypedNamedLocation: TNLRefuel}))
	job.Tasks = append(job.Tasks, TaskPlanners[TaskRefuel](orchestration, TaskContext{}))

	//Go home
	job.Tasks = append(job.Tasks, TaskPlanners[TaskTravelToX](orchestration, TaskContext{TypedNamedLocation: TNLGarageEntrance}))
	job.Tasks = append(job.Tasks, TaskPlanners[TaskPark](orchestration, TaskContext{TypedNamedLocation: TNLParking}))
	job.Tasks = append(job.Tasks, TaskPlanners[TaskTurnToDirection](orchestration, TaskContext{Direction: turtle.parkingDirection}))

}
