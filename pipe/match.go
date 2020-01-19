package pipe

import (
	"fmt"
	"regexp"

	g "github.com/AllenDang/giu"
)

type MatchPipe struct {
	matchWith string
}

func NewMatchPipe() Pipe {
	return &MatchPipe{}
}

func (m *MatchPipe) GetName() string {
	return "M"
}

func (m *MatchPipe) GetTip() string {
	return fmt.Sprintf("Match input string array with regex %s", m.matchWith)
}

func (m *MatchPipe) GetInputType() DataType {
	return DataTypeStringArray
}

func (m *MatchPipe) GetOutputType() DataType {
	return DataTypeStringArray
}

func (m *MatchPipe) GetConfigUI(changed func()) g.Layout {
	return g.Layout{
		g.InputTextV("Match with", 100, &(m.matchWith), 0, nil, changed),
	}
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
