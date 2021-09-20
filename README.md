# go-mlflow
go-mlfow is an unofficial client for MLflow. MLflow is an open source platform for managing the end-to-end machine learning lifecycle. 
The package provides functionality related to MLflow tracking such as list experiments, get experiments by name/id, get runs and search runs.

Install
=======
Since this is a client package for MLflow server, Please ensure that MLflow is up and running.
If MLflow is not installed, you can check the installation instructions [here](https://mlflow.org/docs/latest/quickstart.html).

To install go-mlflow, simply run the following command
```
go get github.com/gagansingh894/go-mlflow
```

Usage
=====
```go
package main

import mlflow "github.com/gagansingh894/go-mlflow/mlflow"

func main() {
    // Create a client for connecting with MLflow server
	m := mlflow.Client("http://localhost", "5000")
    // Create experiment with name 'test1' and location ''
	_ = m.CreateExperiment("test1", "./mlruns/1")
	_ = m.CreateExperiment("test2", "./mlruns/2")
	_ = m.ListExperiments()
}

```
```
# m.CreateExperiment("test1", "./mlruns/1") Output
{
        "experiment_id": "1",
        "error_code": "",
        "message": ""
}

# m.CreateExperiment("test2", "./mlruns/2") Output
{
        "experiment_id": "2",
        "error_code": "",
        "message": ""
}

# m.ListExperiments() Output
{
        "experiments": [
                {
                        "experiment_id": "2",
                        "name": "test2",
                        "artifact_location": "./mlruns/2",
                        "lifecycle_stage": "active"
                },
                {
                        "experiment_id": "1",
                        "name": "test1",
                        "artifact_location": "./mlruns/1",
                        "lifecycle_stage": "active"
                },
                {
                        "experiment_id": "0",
                        "name": "Default",
                        "artifact_location": "./mlruns/0",
                        "lifecycle_stage": "active"
                }
        ]
}

```

