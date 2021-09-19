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

// mlflowapi is an interface which declares the various methods for interacting with MLFLOW server
type mlflowapi interface {
	CreateExperiment(name string, artifactLocation string) CreateExperimentResponse
	ListExperiments() ListExperimentsResponse
	GetExperiment(experimentID string) GetExperimentResponse
	GetExperimentByName(experimentName string) GetExperimentByNameResponse
	DeleteExperiment(experimentID string)
	GetRun(runID string, runUUID string) GetRunResponse
	SearchRun(experimentIDS []string, filter string, runViewType ViewType, maxResults int32, orderBy []string, pageToken string) SearchRunsResponse
}

// Client function creates a new instance of Mlflow struct. The user will use this function to create the client for MLFLOW server
func Client(host string, port string) Mlflow {
	url := host + ":" + port
	client := &http.Client{}
	return Mlflow{host: host, port: port, url: url, client: *client}
}

// prettyPrint is helper function for beautifying the json responses from the MLFLOW server
func prettyPrint(i interface{}) {
	s, _ := json.MarshalIndent(i, "", "\t")
	fmt.Println(string(s))
}

// CreateExperiment is an interface method used to create a new experiment in MLFLOW
func (m *Mlflow) CreateExperiment(name string, artifactLocation string) CreateExperimentResponse {
	endpoint := m.url + "/api/2.0/mlflow/experiments/create"
	requestStruct := CreateExperimentRequest{Name: name, ArtifactLocation: artifactLocation}
	var responseStruct CreateExperimentResponse

	requestJSON, err := json.Marshal(requestStruct)
	if err != nil {
		log.Fatal(err)
	}

	requestBytes := new(bytes.Buffer)
	_ = json.NewEncoder(requestBytes).Encode(string(requestJSON))

	request, _ := http.NewRequest("POST", endpoint, bytes.NewReader(requestBytes.Bytes()))
	response, err := m.client.Do(request)

	if err != nil {
		log.Fatal(err)
	}

	if response.StatusCode == http.StatusOK {
		responseBodyBytes, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		}
		_ = json.Unmarshal(responseBodyBytes, &responseStruct)

	}

	if response.StatusCode == http.StatusBadRequest {
		responseBodyBytes, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		}
		_ = json.Unmarshal(responseBodyBytes, &responseStruct)
	}

	prettyPrint(responseStruct)
	return responseStruct

}

func (m *Mlflow) ListExperiments() ListExperimentsResponse {
	endpoint := m.url + "/api/2.0/mlflow/experiments/list"
	var responseStruct ListExperimentsResponse

	request, _ := http.NewRequest("GET", endpoint, nil)
	response, err := m.client.Do(request)

	if err != nil {
		log.Fatal(err)
	}

	if response.StatusCode == http.StatusOK {
		responseBodyBytes, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		}
		_ = json.Unmarshal(responseBodyBytes, &responseStruct)
	}

	if response.StatusCode == http.StatusBadRequest {
		responseBodyBytes, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		}
		_ = json.Unmarshal(responseBodyBytes, &responseStruct)
	}

	prettyPrint(responseStruct)
	return responseStruct

}

func (m *Mlflow) GetExperiment(experimentID string) GetExperimentResponse {
	endpoint := m.url + "/api/2.0/mlflow/experiments/get"
	requestStruct := GetExperimentRequest{ExperimentID: experimentID}
	var responseStruct GetExperimentResponse

	requestJSON, err := json.Marshal(requestStruct)
	if err != nil {
		log.Fatal(err)
	}

	requestBytes := new(bytes.Buffer)
	_ = json.NewEncoder(requestBytes).Encode(string(requestJSON))

	request, _ := http.NewRequest("GET", endpoint, bytes.NewReader(requestBytes.Bytes()))
	response, err := m.client.Do(request)

	if err != nil {
		log.Fatal(err)
	}

	if response.StatusCode == http.StatusOK {
		responseBodyBytes, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		}
		_ = json.Unmarshal(responseBodyBytes, &responseStruct)
	}

	if response.StatusCode == http.StatusBadRequest {
		responseBodyBytes, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		}
		_ = json.Unmarshal(responseBodyBytes, &responseStruct)
	}

	prettyPrint(responseStruct)
	return responseStruct

}


func (m *Mlflow) GetExperimentByName(experimentName string) GetExperimentByNameResponse {
	endpoint := m.url + "/api/2.0/mlflow/experiments/get-by-name"
	requestStruct := GetExperimentByNameRequest{ExperimentName: experimentName}
	var responseStruct GetExperimentByNameResponse

	requestJSON, err := json.Marshal(requestStruct)
	if err != nil {
		log.Fatal(err)
	}

	requestBytes := new(bytes.Buffer)
	_ = json.NewEncoder(requestBytes).Encode(string(requestJSON))

	request, _ := http.NewRequest("GET", endpoint, bytes.NewReader(requestBytes.Bytes()))
	response, err := m.client.Do(request)

	if err != nil {
		log.Fatal(err)
	}

	if response.StatusCode == http.StatusOK {
		responseBodyBytes, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		}
		_ = json.Unmarshal(responseBodyBytes, &responseStruct)
	}

	if response.StatusCode == http.StatusBadRequest {
		responseBodyBytes, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		}
		_ = json.Unmarshal(responseBodyBytes, &responseStruct)
	}

	prettyPrint(responseStruct)
	return responseStruct

}

func (m *Mlflow) DeleteExperiment(experimentID string) {
	endpoint := m.url + "/api/2.0/mlflow/experiments/delete"
	requestStruct := DeleteExperimentRequest{ExperimentID: experimentID}

	requestJSON, err := json.Marshal(requestStruct)
	if err != nil {
		log.Fatal(err)
	}

	requestBytes := new(bytes.Buffer)
	_ = json.NewEncoder(requestBytes).Encode(string(requestJSON))

	request, _ := http.NewRequest("POST", endpoint, bytes.NewReader(requestBytes.Bytes()))
	response, err := m.client.Do(request)

	if err != nil {
		log.Fatal(err)
	}

	if response.StatusCode == http.StatusOK {
		log.Println("OK")

	}

	if response.StatusCode == http.StatusBadRequest {
		log.Println("BadRequest")
	}


}

func (m *Mlflow) GetRun(runID string, runUUID string) GetRunResponse {
	endpoint := m.url + "/api/2.0/mlflow/runs/get"
	requestStruct := GetRunRequest{RunID: runID, RunUUID: runUUID}
	var responseStruct GetRunResponse

	requestJSON, err := json.Marshal(requestStruct)
	if err != nil {
		log.Fatal(err)
	}

	requestBytes := new(bytes.Buffer)
	_ = json.NewEncoder(requestBytes).Encode(string(requestJSON))

	request, _ := http.NewRequest("POST", endpoint, bytes.NewReader(requestBytes.Bytes()))
	response, err := m.client.Do(request)

	if err != nil {
		log.Fatal(err)
	}

	if response.StatusCode == http.StatusOK {
		responseBodyBytes, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		}
		_ = json.Unmarshal(responseBodyBytes, &responseStruct)
	}

	if response.StatusCode == http.StatusBadRequest {
		responseBodyBytes, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		}
		_ = json.Unmarshal(responseBodyBytes, &responseStruct)
	}

	prettyPrint(responseStruct)
	return responseStruct

}

func (m *Mlflow) SearchRun(experimentIDS []string, filter string, runViewType ViewType, maxResults int32, orderBy []string, pageToken string) SearchRunsResponse {
	endpoint := m.url + "/api/2.0/mlflow/runs/search"
	requestStruct := SearchRunsRequest{ExperimentIDS: experimentIDS, Filter: filter, RunViewType: runViewType, MaxResults: maxResults, OrderBy: orderBy, PageToken: pageToken}
	var responseStruct SearchRunsResponse

	requestJSON, err := json.Marshal(requestStruct)
	if err != nil {
		log.Fatal(err)
	}

	requestBytes := new(bytes.Buffer)
	_ = json.NewEncoder(requestBytes).Encode(string(requestJSON))

	request, _ := http.NewRequest("POST", endpoint, bytes.NewReader(requestBytes.Bytes()))
	response, err := m.client.Do(request)

	if err != nil {
		log.Fatal(err)
	}

	if response.StatusCode == http.StatusOK {
		responseBodyBytes, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		}
		_ = json.Unmarshal(responseBodyBytes, &responseStruct)
	}

	if response.StatusCode == http.StatusBadRequest {
		responseBodyBytes, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		}
		_ = json.Unmarshal(responseBodyBytes, &responseStruct)
	}

	prettyPrint(responseStruct)
	return responseStruct
}