package pipe

import (
	"strings"

	g "github.com/AllenDang/giu"
)

type TrimPipe struct {
	trimWith string
}

func NewTrimPipe() Pipe {
	return &TrimPipe{}
}

func (t *TrimPipe) GetName() string {
	return "T"
}

func (t *TrimPipe) GetTip() string {
	return "Trim returns a slice of the string s with all leading and trailing Unicode code points contained in cutset removed"
}

func (t *TrimPipe) GetInputType() DataType {
	return DataTypeStringArray
}

func (t *TrimPipe) GetOutputType() DataType {
	return DataTypeStringArray
}

func (t *TrimPipe) GetConfigUI(changed func()) g.Layout {
	return g.Layout{
		g.InputTextV("Trim with", 100, &(t.trimWith), 0, nil, changed),
	}
}

func (t *TrimPipe) Process(data interface{}) interface{} {
	if strs, ok := data.([]string); ok {
		var results []string
		for _, s := range strs {
			results = append(results, strings.Trim(s, t.trimWith))
		}

		return results
	}

	return []string{"Error: Trim only accepts string array as input"}
}
