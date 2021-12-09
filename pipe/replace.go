package pipe

import (
	"encoding/gob"
	"fmt"
	"regexp"

	g "github.com/AllenDang/giu"
)

type ReplacePipe struct {
	Replace string
	With    string
}

func init() {
	gob.Register(&ReplacePipe{})
}

func NewReplacePipe() Pipe {
	return &ReplacePipe{}
}

func (r *ReplacePipe) GetName() string {
	return "R"
}

func (r *ReplacePipe) GetTip() string {
	return fmt.Sprintf("Replace each string of input string array from %s to %s", r.Replace, r.With)
}

func (r *ReplacePipe) GetInputType() DataType {
	return DataTypeStringArray
}

func (r *ReplacePipe) GetOutputType() DataType {
	return DataTypeStringArray
}

func (r *ReplacePipe) GetConfigUI(changed func()) g.Layout {
	return g.Layout{
		g.InputText(&(r.Replace)).Label("Replace").Size(100).OnChange(changed),
		g.InputText(&(r.With)).Label("With").Size(100).OnChange(changed),
	}
}

func (r *ReplacePipe) Process(data interface{}) interface{} {
	if strs, ok := data.([]string); ok {
		re, err := regexp.Compile(r.Replace)
		if err != nil {
			return []string{"Error: Invalid regex"}
		}

		var result []string
		for _, s := range strs {
			result = append(result, re.ReplaceAllString(s, r.With))
		}

		return result
	}

	return []string{"Error: Replace only accepts string array as input"}
}
