package pipe

import (
	"encoding/gob"

	g "github.com/AllenDang/giu"
)

type Column2RowPipe struct{}

func init() {
	gob.Register(&Column2RowPipe{})
}

func NewColumn2RowPipe() Pipe {
	return &Column2RowPipe{}
}

func (c *Column2RowPipe) GetName() string {
	return "C2R"
}

func (c *Column2RowPipe) GetTip() string {
	return "Shift table's column to row"
}

func (c *Column2RowPipe) GetInputType() DataType {
	return DataTypeTable
}

func (c *Column2RowPipe) GetOutputType() DataType {
	return DataTypeTable
}

func (c *Column2RowPipe) GetConfigUI(changed func()) g.Layout {
	return nil
}

func (c *Column2RowPipe) Process(data interface{}) interface{} {
	if table, ok := data.([][]string); ok {
		if len(table) > 0 && len(table[0]) > 0 {
			columnCount := len(table[0])
			rows := make([][]string, columnCount)

			for _, r := range table {
				for i, c := range r {
					if i < columnCount {
						rows[i] = append(rows[i], c)
					}
				}
			}

			return rows
		}
	}

	return [][]string{[]string{"Error: Column2RowPipe only accepts table as input"}}
}
