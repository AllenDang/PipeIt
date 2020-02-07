package pipe

import (
	"encoding/gob"
	"strings"

	g "github.com/AllenDang/giu"
)

type FieldsPipe struct{}

func init() {
	gob.Register(&FieldsPipe{})
}

func NewFieldsPipe() Pipe {
	return &FieldsPipe{}
}

func (f *FieldsPipe) GetName() string {
	return "F"
}

func (f *FieldsPipe) GetTip() string {
	return "Fields splits the string s around each instance of one or more consecutive white space characters"
}

func (f *FieldsPipe) GetInputType() DataType {
	return DataTypeString
}

func (f *FieldsPipe) GetOutputType() DataType {
	return DataTypeStringArray
}

func (f *FieldsPipe) GetConfigUI(changed func()) g.Layout {
	return nil
}

func (f *FieldsPipe) Process(data interface{}) interface{} {
	if str, ok := data.(string); ok {
		return strings.Fields(str)
	}

	return []string{"Error: Fields only accepts string as input"}
}
