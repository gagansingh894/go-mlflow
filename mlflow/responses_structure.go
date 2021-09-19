package mlflow

type CreateExperimentResponse struct {
	ExperimentID string `json:"experiment_id"`
	ErrorCode    string `json:"error_code"`
	Message      string `json:"message"`
}

type ListExperimentsResponse struct {
	Experiments []Experiment `json:"experiments"`
}

type GetExperimentResponse struct {
	Experiment Experiment `json:"experiment"`
}

type GetExperimentByNameResponse struct {
	Experiment Experiment `json:"experiment"`
}

type GetRunResponse struct {
	Run Run `json:"run"`
}

type SearchRunsResponse struct {
	Runs          []Run  `json:"runs"`
	NextPageToken string `json:"next_page_token"`
}
