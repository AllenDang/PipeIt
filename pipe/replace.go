package pipe

import (
	"fmt"
	"strings"

	g "github.com/AllenDang/giu"
)

type ReplacePipe struct {
	replace string
	with    string
}

func NewReplacePipe() Pipe {
	return &ReplacePipe{}
}

func (r *ReplacePipe) GetName() string {
	return "R"
}

func (r *ReplacePipe) GetTip() string {
	return fmt.Sprintf("Replace each string of input string array from %s to %s", r.replace, r.with)
}

func (r *ReplacePipe) GetInputType() DataType {
	return DataTypeStringArray
}

func (r *ReplacePipe) GetOutputType() DataType {
	return DataTypeStringArray
}

func (r *ReplacePipe) GetConfigUI(changed func()) g.Layout {
	return g.Layout{
		g.InputTextV("Replace", 100, &(r.replace), 0, nil, changed),
		g.InputTextV("With", 100, &(r.with), 0, nil, changed),
	}
}

func (r *ReplacePipe) Process(data interface{}) interface{} {
	if strs, ok := data.([]string); ok {
		var result []string
		for _, s := range strs {
			result = append(result, strings.ReplaceAll(s, r.replace, r.with))
		}

		return result
	}

	return []string{"Error: Replace only accepts string array as input"}
}
