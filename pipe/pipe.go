package pipe

import (
	"bytes"
	"encoding/gob"
	"io"

	g "github.com/AllenDang/giu"
)

type DataType int

const (
	DataTypeString DataType = iota
	DataTypeStringArray
	DataTypeTable
)

type Parameter struct {
	Type  DataType
	Value interface{}
}

type Pipe interface {
	// Get information for the pipe including name, bgColor, nameColor and borderColor
	GetName() string
	GetTip() string
	GetInputType() DataType
	GetOutputType() DataType
	GetConfigUI(changed func()) g.Layout
	Process(data interface{}) interface{}
}

type Pipeline []Pipe

func EncodePipeline(pl Pipeline) (*bytes.Buffer, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(&pl)
	if err != nil {
		return nil, err
	}

	return &buf, nil
}

func DecodePipeline(r io.Reader) (*Pipeline, error) {
	var pl Pipeline
	dec := gob.NewDecoder(r)
	err := dec.Decode(&pl)
	if err != nil {
		return nil, err
	}

	return &pl, nil
}
