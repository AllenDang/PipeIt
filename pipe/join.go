package pipe

import (
	"encoding/gob"
	"fmt"
	"strings"

	g "github.com/AllenDang/giu"
)

type JoinPipe struct {
	JoinWith string
}

func init() {
	gob.Register(&JoinPipe{})
}

func NewJoinPipe() Pipe {
	return &JoinPipe{
		JoinWith: ",",
	}
}

func (j *JoinPipe) GetName() string {
	return "J"
}

func (j *JoinPipe) GetTip() string {
	return fmt.Sprintf("Join string array with %s", j.JoinWith)
}

func (j *JoinPipe) GetInputType() DataType {
	return DataTypeStringArray
}

func (j *JoinPipe) GetOutputType() DataType {
	return DataTypeString
}

func (j *JoinPipe) GetConfigUI(changed func()) g.Layout {
	return g.Layout{
		g.InputTextV("Join With", 100, &(j.JoinWith), 0, nil, changed),
	}
}

func (j *JoinPipe) Process(data interface{}) interface{} {
	if strs, ok := data.([]string); ok {
		return strings.Join(strs, j.JoinWith)
	}

	return "Error: Join only accepts string array as input"
}
