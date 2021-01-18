package domain

type Orchestration struct {
	Name              string            `json:"name"`
	OrchestrationType OrchestrationType `json:"orchestrationType"`
	FeaturesRequired  []TurtleFeature   `json:"featuresRequired"`
	Priority          int               `json:"priority"`
	JobType           JobType           `json:"jobType"`
	Spec              OrchestrationSpec `json:"spec"`
}

type OrchestrationType string

const (

	/*
		Cron: a turtle is allocation at each interval `everySeconds`
	*/
	OrchestrationTypeCron OrchestrationType = "ORCH_CRON"

	/*
		Deployment: a new worker is sent after the work is done
	*/
	OrchestrationTypeDeployment OrchestrationType = "ORCH_DEPLOYMENT"
)

type OrchestrationSpec struct {

	//Cron specific
	EverySeconds  int           `json:"everySeconds"`
	StartLocation NamedLocation `json:"startLocation"`

	//Deployment specific
	MinimumFuel int `json:"minimumFuel"`

	//Global
	StartDirection Direction         `json:"startDirection"`
	Meta           map[string]string `json:"meta"`
	MetaIntegers   map[string]int    `json:"metaIntegers"`
}
