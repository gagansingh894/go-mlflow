package mlflow

type CreateExperimentRequest struct {
	Name             string `json:"name"`
	ArtifactLocation string `json:"artifact_location"`
}

type GetExperimentRequest struct {
	ExperimentID string `json:"experiment_id"`
}

type GetExperimentByNameRequest struct {
	ExperimentName string `json:"experiment_name"`
}

type DeleteExperimentRequest struct {
	ExperimentID string `json:"experiment_id"`
}

type GetRunRequest struct {
	RunID   string `json:"run_id"`
	RunUUID string `json:"run_uuid"`
}

type SearchRunsRequest struct {
	ExperimentIDS []string `json:"experiment_ids"`
	Filter        string   `json:"filter"`
	RunViewType   ViewType `json:"run_view_type"`
	MaxResults    int32    `json:"max_results"`
	OrderBy       []string `json:"order_by"`
	PageToken     string   `json:"page_token"`
}
