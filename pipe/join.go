package pipe

import (
	"fmt"
	"strings"
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
	return fmt.Sprintf("Join(%s)", j.joinWith)
}

func (j *JoinPipe) GetInputType() DataType {
	return DataTypeStringArray
}

func (j *JoinPipe) GetOutputType() DataType {
	return DataTypeString
}

func (j *JoinPipe) GetParameters() map[string]*Parameter {
	params := make(map[string]*Parameter)
	params["JoinWith"] = &Parameter{Type: DataTypeString, Value: &(j.joinWith)}
	return params
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
