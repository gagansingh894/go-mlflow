package mlflow

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Mlflow is the base struct which is used to interact with the MLFLOW server
type Mlflow struct {
	host   string
	port   string
	url    string
	client http.Client
}

type parameter struct {
	name             string
	artifactLocation string
	experimentID     string
	experimentName   string
	runID            string
	runUUID          string
	experimentIDS    []string
	filter           string
	runViewType      ViewType
	maxResults       int32
	orderBy          []string
	pageToken        string
}

// mlflowapi is an interface which declares the various methods for interacting with MLFLOW server
type Mlflowapi interface {
	CreateExperiment(name string, artifactLocation string) CreateExperimentResponse
	ListExperiments() ListExperimentsResponse
	GetExperiment(experimentID string) GetExperimentResponse
	GetExperimentByName(experimentName string) GetExperimentByNameResponse
	DeleteExperiment(experimentID string)
	GetRun(runID string, runUUID string) GetRunResponse
	SearchRun(experimentIDS []string, filter string, runViewType ViewType, maxResults int32, orderBy []string, pageToken string) SearchRunsResponse
}

// NewClient function creates a new instance of Mlflow struct. The user will use this function to create the client for MLFLOW server
func NewClient(host string, port string) *Mlflow {
	url := host + ":" + port
	client := &http.Client{}
	mlflow := Mlflow{host: host, port: port, url: url, client: *client}
	return &mlflow
}

// prettyPrint is helper function for beautifying the json responses from the MLFLOW server
func prettyPrint(i interface{}) {
	s, _ := json.MarshalIndent(i, "", "\t")
	fmt.Println(string(s))
}

func base(m *Mlflow, endpoint string, requestStruct interface{}, requestType string) []byte {
	var defaultBytes []byte
	requestJSON, err := json.Marshal(requestStruct)
	if err != nil {
		log.Fatal(err)
	}

	requestBytes := new(bytes.Buffer)
	_ = json.NewEncoder(requestBytes).Encode(string(requestJSON))

	request, _ := http.NewRequest(requestType, endpoint, bytes.NewReader(requestBytes.Bytes()))
	response, err := m.client.Do(request)

	if err != nil {
		log.Fatal(err)
	}

	if response.StatusCode == http.StatusOK {
		responseBodyBytes, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		} else {
			return responseBodyBytes
		}
		//_ = json.Unmarshal(responseBodyBytes, &responseStruct)

	} else if response.StatusCode == http.StatusBadRequest {
		responseBodyBytes, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		} else {
			return responseBodyBytes
		}
		//_ = json.Unmarshal(responseBodyBytes, &responseStruct)
	} else {
		log.Fatal("Unexpected error occured")
	}
	return defaultBytes
}

// CreateExperiment is an interface method used to create a new experiment in MLFLOW
func (m *Mlflow) CreateExperiment(name string, artifactLocation string) CreateExperimentResponse {
	endpoint := m.url + "/api/2.0/mlflow/experiments/create"
	requestStruct := CreateExperimentRequest{Name: name, ArtifactLocation: artifactLocation}
	var responseStruct CreateExperimentResponse
	baseResponse := base(m, endpoint, requestStruct, "POST")
	_ = json.Unmarshal(baseResponse, &responseStruct)
	prettyPrint(responseStruct)
	return responseStruct
}

func (m *Mlflow) ListExperiments() ListExperimentsResponse {
	endpoint := m.url + "/api/2.0/mlflow/experiments/list"
	var responseStruct ListExperimentsResponse
	baseResponse := base(m, endpoint, nil, "GET")
	_ = json.Unmarshal(baseResponse, &responseStruct)
	prettyPrint(responseStruct)
	return responseStruct
}

func (m *Mlflow) GetExperiment(experimentID string) GetExperimentResponse {
	endpoint := m.url + "/api/2.0/mlflow/experiments/get"
	requestStruct := GetExperimentRequest{ExperimentID: experimentID}
	var responseStruct GetExperimentResponse
	baseResponse := base(m, endpoint, requestStruct, "GET")
	_ = json.Unmarshal(baseResponse, &responseStruct)
	prettyPrint(responseStruct)
	return responseStruct
}

func (m *Mlflow) GetExperimentByName(experimentName string) GetExperimentByNameResponse {
	endpoint := m.url + "/api/2.0/mlflow/experiments/get-by-name"
	requestStruct := GetExperimentByNameRequest{ExperimentName: experimentName}
	var responseStruct GetExperimentByNameResponse
	baseResponse := base(m, endpoint, requestStruct, "GET")
	_ = json.Unmarshal(baseResponse, &responseStruct)
	prettyPrint(responseStruct)
	return responseStruct
}

func (m *Mlflow) DeleteExperiment(experimentID string) {
	endpoint := m.url + "/api/2.0/mlflow/experiments/delete"
	requestStruct := DeleteExperimentRequest{ExperimentID: experimentID}
	_ = base(m, endpoint, requestStruct, "GET")
}

func (m *Mlflow) GetRun(runID string, runUUID string) GetRunResponse {
	endpoint := m.url + "/api/2.0/mlflow/runs/get"
	requestStruct := GetRunRequest{RunID: runID, RunUUID: runUUID}
	var responseStruct GetRunResponse
	baseResponse := base(m, endpoint, requestStruct, "POST")
	_ = json.Unmarshal(baseResponse, &responseStruct)
	prettyPrint(responseStruct)
	return responseStruct
}

func (m *Mlflow) SearchRun(experimentIDS []string, filter string, runViewType ViewType, maxResults int32, orderBy []string, pageToken string) SearchRunsResponse {
	endpoint := m.url + "/api/2.0/mlflow/runs/search"
	requestStruct := SearchRunsRequest{ExperimentIDS: experimentIDS, Filter: filter, RunViewType: runViewType, MaxResults: maxResults, OrderBy: orderBy, PageToken: pageToken}
	var responseStruct SearchRunsResponse
	baseResponse := base(m, endpoint, requestStruct, "POST")
	_ = json.Unmarshal(baseResponse, &responseStruct)
	prettyPrint(responseStruct)
	return responseStruct
}
