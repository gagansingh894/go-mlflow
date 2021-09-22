package main

import (
	mlflow "github.com/gagansingh894/go-mlflow/mlflow"
	"log"
)

func main() {
	// Create a client for connecting with MLflow server
	m := mlflow.NewClient("http://localhost", "5000")
	// Create experiment with name 'test1' and location './mlruns/1'
	_, err := m.CreateExperiment("test1", "./mlruns/1")
	if err != nil {
		log.Fatal(err)
	}
	// Create experiment with name 'test2' and location './mlruns/2'
	_, err = m.CreateExperiment("test2", "./mlruns/2")
	if err != nil {
		log.Fatal(err)
	}
	// List Experiments
	_, err = m.ListExperiments()
	if err != nil {
		log.Fatal(err)
	}
}
