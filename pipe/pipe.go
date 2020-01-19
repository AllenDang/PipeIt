package pipe

import (
	g "github.com/AllenDang/giu"
)

type DataType int

const (
	DataTypeString DataType = iota
	DataTypeInt
	DataTypeStringArray
)

type Parameter struct {
	Type  DataType
	Value interface{}
}

type Pipe interface {
	GetName() string
	GetInputType() DataType
	GetOutputType() DataType
	GetConfigUI(changed func()) g.Layout
	Process(data interface{}) interface{}
}

type Pipeline []Pipe
