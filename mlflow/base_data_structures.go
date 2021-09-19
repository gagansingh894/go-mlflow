package mlflow

type Experiment struct {
	ExperimentID     string `json:"experiment_id"`
	Name             string `json:"name"`
	ArtifactLocation string `json:"artifact_location"`
	LifecycleStage   string `json:"lifecycle_stage"`
}

type Experiments struct {
	Experiments []Experiment `json:"experiments"`
}

type ExperimentTag struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type Run struct {
	Info RunInfo `json:"info"`
	Data RunData `json:"data"`
}

type RunInfo struct {
	RunID          string    `json:"run_id"`
	RunUUID        string    `json:"run_uuid"`
	ExperimentID   string    `json:"experiment_id"`
	UserID         string    `json:"user_id"`
	StartTime      int64     `json:"start_time"`
	Status         RunStatus `json:"status"`
	EndTime        int64     `json:"end_time"`
	ArtifactURI    string    `json:"artifact_uri"`
	LifecycleStage string    `json:"lifecycle_stage"`
}

type RunData struct {
	Metrics []Metric `json:"metrics"`
	Params  []Param  `json:"params"`
	Tags    []RunTag `json:"tags"`
}

type Metric struct {
	Key   string  `json:"key"`
	Value float64 `json:"value"`
}

type Param struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type RunTag struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type ModelVersionStatus string

const (
	PENDING_REGISTRATION ModelVersionStatus = "PENDING_REGISTRATION"
	FAILED_REGISTRATION  ModelVersionStatus = "FAILED_REGISTRATION"
	READY                ModelVersionStatus = "READY"
)

type RunStatus string

const (
	RUNNING   RunStatus = "RUNNING"
	SCHEDULED RunStatus = "SCHEDULED"
	FINISHED  RunStatus = "FINISHED"
	FAILED    RunStatus = "FAILED"
	KILLED    RunStatus = "KILLED"
)

type ViewType string

const (
	ACTIVE_ONLY  ViewType = "ACTIVE_ONLY"
	DELETED_ONLY ViewType = "DELETED_ONLY"
	ALL          ViewType = "ALL"
)
