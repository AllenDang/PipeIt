package pipe

import (
	"encoding/gob"
	"fmt"
	"regexp"

	g "github.com/AllenDang/giu"
)

type MatchPipe struct {
	MatchWith string
}

func init() {
	gob.Register(&MatchPipe{})
}

func NewMatchPipe() Pipe {
	return &MatchPipe{}
}

func (m *MatchPipe) GetName() string {
	return "M"
}

func (m *MatchPipe) GetTip() string {
	return fmt.Sprintf("Match input string array with regex %s", m.MatchWith)
}

func (m *MatchPipe) GetInputType() DataType {
	return DataTypeStringArray
}

func (m *MatchPipe) GetOutputType() DataType {
	return DataTypeStringArray
}

func (m *MatchPipe) GetConfigUI(changed func()) g.Layout {
	return g.Layout{
		g.InputText(&(m.MatchWith)).Label("Match with").Size(100).OnChange(changed),
	}
}

func (m *MatchPipe) Process(data interface{}) interface{} {
	if strs, ok := data.([]string); ok {
		var result []string
		for _, s := range strs {
			if matched, _ := regexp.MatchString(m.MatchWith, s); matched {
				result = append(result, s)
			}
		}

		return result
	}

	return []string{"Error: Match only accepts string array as input"}
}
