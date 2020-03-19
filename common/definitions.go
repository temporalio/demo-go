package common

const (
	// Domain used by samples for hosting workflows
	Domain = "samples"
	// WorkflowTaskList is the queue used by worker to pull workflow tasks
	WorkflowTaskList = "samples_workflow_tl"
	// ActivityTaskList is the queue used by worker to pull activity tasks
	ActivityTaskList = "samples_activity_tl"
	// Host is the host/port used by client to connect to temporal server
	Host = "127.0.0.1:7233"
)
