package pipe

import (
	"fmt"
	"strings"
)

type SplitPipe struct {
	splitWith string
}

func NewSplitPipe() Pipe {
	return &SplitPipe{
		splitWith: ",",
	}
}

func (s *SplitPipe) GetName() string {
	return fmt.Sprintf("Split(%s)", s.splitWith)
}

func (s *SplitPipe) GetInputType() DataType {
	return DataTypeString
}

func (s *SplitPipe) GetOutputType() DataType {
	return DataTypeStringArray
}

func (s *SplitPipe) GetParameters() map[string]*Parameter {
	params := make(map[string]*Parameter)
	params["SplitWith"] = &Parameter{Type: DataTypeString, Value: &(s.splitWith)}
	return params
}

func (s *SplitPipe) Process(data interface{}) interface{} {
	if str, ok := data.(string); ok {
		return strings.Split(str, s.splitWith)
	}

	return []string{"Error: Split only accept string as input"}
}
