package pipe

import (
	"fmt"
	"strings"

	g "github.com/AllenDang/giu"
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

func (s *SplitPipe) GetConfigUI(changed func()) g.Layout {
	return g.Layout{
		g.InputTextV("Split with", 100, &(s.splitWith), 0, nil, changed),
	}
}

func (s *SplitPipe) Process(data interface{}) interface{} {
	if str, ok := data.(string); ok {
		return strings.Split(str, s.splitWith)
	}

	return []string{"Error: Split only accept string as input"}
}
