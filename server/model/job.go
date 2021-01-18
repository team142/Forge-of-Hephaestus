package model

/*
	Overarching job is a particular service carried out by Turtles
	A consists of a list of tasks. These are determined by its dynamic target
*/
type Job struct {
	Name          JobType
	Tasks         []Task
	CurrentTaskID string
}

func (j Job) IsActualWorkDone() bool {
	actualWorkDone := false
	for _, t := range j.Tasks {
		if t.ID == j.CurrentTaskID {
			break
		}
		if t.ActualWork {
			actualWorkDone = true
		}
	}
	return actualWorkDone
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
