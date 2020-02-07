package pipe

import (
	"encoding/gob"
	"fmt"
	"strings"

	g "github.com/AllenDang/giu"
)

type LinePipe struct{}

func init() {
	gob.Register(&LinePipe{})
}

func NewLinePipe() Pipe {
	return &LinePipe{}
}

func (l *LinePipe) GetName() string {
	return "L"
}

func (l *LinePipe) GetTip() string {
	return "Output input string line by line"
}

func (l *LinePipe) GetInputType() DataType {
	return DataTypeStringArray
}

func (l *LinePipe) GetOutputType() DataType {
	return DataTypeString
}

func (l *LinePipe) GetConfigUI(changed func()) g.Layout {
	return nil
}

func (l *LinePipe) Process(data interface{}) interface{} {
	if strs, ok := data.([]string); ok {
		var sb strings.Builder
		for _, s := range strs {
			sb.WriteString(fmt.Sprintf("%s\n", s))
		}

		return sb.String()
	}

	return "Error: Line only accepts string array as input"
}
