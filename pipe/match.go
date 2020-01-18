package pipe

import (
	"fmt"
	"regexp"
)

type MatchPipe struct {
	matchWith string
}

func NewMatchPipe() Pipe {
	return &MatchPipe{}
}

func (m *MatchPipe) GetName() string {
	return fmt.Sprintf("Match(%s)", m.matchWith)
}

func (m *MatchPipe) GetInputType() DataType {
	return DataTypeStringArray
}

func (m *MatchPipe) GetOutputType() DataType {
	return DataTypeStringArray
}

func (m *MatchPipe) GetParameters() map[string]*Parameter {
	params := make(map[string]*Parameter)
	params["MatchWith"] = &Parameter{Type: DataTypeString, Value: &(m.matchWith)}
	return params
}

func (m *MatchPipe) Process(data interface{}) interface{} {
	if strs, ok := data.([]string); ok {
		var result []string
		for _, s := range strs {
			if matched, _ := regexp.MatchString(m.matchWith, s); matched {
				result = append(result, s)
			}
		}

		return result
	}

	return []string{"Error: Match only accepts string array as input"}
}
