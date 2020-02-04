package pipe

import (
	"fmt"
	"strings"

	g "github.com/AllenDang/giu"
)

type FmtRowPipe struct {
	fmtStr string
}

func NewFmtRowPipe() Pipe {
	return &FmtRowPipe{
		fmtStr: "",
	}
}

func (f *FmtRowPipe) GetName() string {
	return "FR"
}

func (f *FmtRowPipe) GetTip() string {
	return "Format row by using %[1]s %[2]s .. %[n]s to reference columns"
}

func (f *FmtRowPipe) GetInputType() DataType {
	return DataTypeTable
}

func (f *FmtRowPipe) GetOutputType() DataType {
	return DataTypeString
}

func (f *FmtRowPipe) GetConfigUI(changed func()) g.Layout {
	return g.Layout{
		g.InputTextV("Fmt string", 300, &(f.fmtStr), 0, nil, changed),
	}
}

func (f *FmtRowPipe) Process(data interface{}) interface{} {
	if table, ok := data.([][]string); ok {
		var sb strings.Builder

		for _, r := range table {
			var tempRow []interface{}
			for _, c := range r {
				tempRow = append(tempRow, c)
			}
			sb.WriteString(fmt.Sprintf(f.fmtStr+"\n", tempRow...))
		}

		return sb.String()
	}

	return "Error: FmtRow only accepts table as input"
}
