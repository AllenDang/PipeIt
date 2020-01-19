package pipe

import (
	"fmt"
	"strings"

	g "github.com/AllenDang/giu"
)

type JoinPipe struct {
	joinWith string
}

func NewJoinPipe() Pipe {
	return &JoinPipe{
		joinWith: ",",
	}
}

func (j *JoinPipe) GetName() string {
	return "J"
}

func (j *JoinPipe) GetTip() string {
	return fmt.Sprintf("Join string array with %s", j.joinWith)
}

func (j *JoinPipe) GetInputType() DataType {
	return DataTypeStringArray
}

func (j *JoinPipe) GetOutputType() DataType {
	return DataTypeString
}

func (j *JoinPipe) GetConfigUI(changed func()) g.Layout {
	return g.Layout{
		g.InputTextV("Join With", 100, &(j.joinWith), 0, nil, changed),
	}
}

func (j *JoinPipe) Process(data interface{}) interface{} {
	if strs, ok := data.([]string); ok {
		result := ""
		for _, s := range strs {
			result += fmt.Sprintf("%s%s", s, j.joinWith)
		}

		return strings.Trim(result, j.joinWith)
	}

	return "Error: Join only accepts string array as input"
}
