package factory

import (
	"bytes"
	"encoding/gob"
)

type Office struct {
	Suite                     int
	StreetName, City, Country string
}

type Worker struct {
	Name   string
	Office *Office
}

func (w *Worker) DeepCopy() *Worker {
	b := bytes.Buffer{}
	e := gob.NewEncoder(&b)
	_ = e.Encode(w)

	result := Worker{}
	d := gob.NewDecoder(&b)
	_ = d.Decode(&result)

	return &result
}

func newWorker(proto *Worker, name string, suite int) *Worker {
	result := proto.DeepCopy()
	result.Name = name
	result.Office.Suite = suite
	return result
}

var mainOfficeWorker = Worker{
	Office: &Office{
		StreetName: "123 Main",
		City:       "Berlin",
		Country:    "DE",
	},
}

var auxOfficeWorker = Worker{
	Office: &Office{
		StreetName: "123 Aux",
		City:       "Berlin",
		Country:    "DE",
	},
}

func NewMainOfficeWorker(name string, suite int) *Worker {
	return newWorker(&mainOfficeWorker, name, suite)
}

func NewAuxOfficeWorker(name string, suite int) *Worker {
	return newWorker(&auxOfficeWorker, name, suite)
}
