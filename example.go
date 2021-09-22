package main

import (
	mlflow "github.com/gagansingh894/go-mlflow/mlflow"
)

func main() {
	// Create a client for connecting with MLflow server
	m := mlflow.NewClient("http://localhost", "5000")
	// Create experiment with name 'test1' and location './mlruns/1'
	_ = m.CreateExperiment("test1", "./mlruns/1")
	// Create experiment with name 'test2' and location './mlruns/2'
	_ = m.CreateExperiment("test2", "./mlruns/2")
	// List Experiments
	_ = m.ListExperiments()
}
