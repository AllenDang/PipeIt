package pipe

import (
	"encoding/gob"
	"regexp"

	g "github.com/AllenDang/giu"
)

type TablePipe struct {
	SplitRowWith    string
	SplitColumnWith string
}

func init() {
	gob.Register(&TablePipe{})
}

func NewTablePipe() Pipe {
	return &TablePipe{
		SplitRowWith:    "\n",
		SplitColumnWith: ",",
	}
}

func (t *TablePipe) GetName() string {
	return "T"
}

func (t *TablePipe) GetTip() string {
	return "Table parse input string to rows and columns"
}

func (t *TablePipe) GetInputType() DataType {
	return DataTypeString
}

func (t *TablePipe) GetOutputType() DataType {
	return DataTypeTable
}

func (t *TablePipe) GetConfigUI(changed func()) g.Layout {
	return g.Layout{
		g.InputTextV("Split row with", 100, &(t.SplitRowWith), 0, nil, changed),
		g.InputTextV("Split column with", 100, &(t.SplitColumnWith), 0, nil, changed),
	}
}

func (t *TablePipe) Process(data interface{}) interface{} {
	if str, ok := data.(string); ok {
		re, err := regexp.Compile(t.SplitRowWith)
		if err != nil {
			return [][]string{[]string{err.Error()}}
		}

		ce, err := regexp.Compile(t.SplitColumnWith)
		if err != nil {
			return [][]string{[]string{err.Error()}}
		}

		tempRows := re.Split(str, -1)

		var rows [][]string

		for _, r := range tempRows {
			rows = append(rows, ce.Split(r, -1))
		}

		return rows
	}

	return [][]string{[]string{"Error: Table only accepts string as input"}}
}
