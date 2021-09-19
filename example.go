package main

import (
	mlflow "www.github.com/gagansingh894/go-mlflow/mlflow"
)

func main() {
	m := mlflow.Client("http://localhost", "5000")
	_ = m.CreateExperiment("test2", "./mlruns/2")
}
