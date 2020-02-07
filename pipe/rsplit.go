package pipe

import (
	"encoding/gob"
	"fmt"
	"regexp"

	g "github.com/AllenDang/giu"
)

type RegexpSplitPipe struct {
	SplitWith string
}

func init() {
	gob.Register(&RegexpSplitPipe{})
}

func NewRegexpSplitPipe() Pipe {
	return &RegexpSplitPipe{
		SplitWith: ",",
	}
}

func (r *RegexpSplitPipe) GetName() string {
	return "S"
}

func (r *RegexpSplitPipe) GetTip() string {
	return fmt.Sprintf("Split input string with regexp: %s", r.SplitWith)
}

func (r *RegexpSplitPipe) GetInputType() DataType {
	return DataTypeString
}

func (r *RegexpSplitPipe) GetOutputType() DataType {
	return DataTypeStringArray
}

func (r *RegexpSplitPipe) GetConfigUI(changed func()) g.Layout {
	return g.Layout{
		g.InputTextV("Split with", 100, &(r.SplitWith), 0, nil, changed),
	}
}

func (r *RegexpSplitPipe) Process(data interface{}) interface{} {
	if str, ok := data.(string); ok {
		re, err := regexp.Compile(r.SplitWith)
		if err == nil {
			return re.Split(str, -1)
		} else {
			return []string{err.Error()}
		}
	}

	return []string{"Error: RegexpSplit only accepts string as input"}
}
