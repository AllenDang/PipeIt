package pipe

type DataType int

const (
	DataTypeString DataType = iota
	DataTypeInt
	DataTypeStringArray
)

type Parameter struct {
	Type  DataType
	Value interface{}
}

type Pipe interface {
	GetName() string
	GetInputType() DataType
	GetOutputType() DataType
	GetParameters() map[string]*Parameter
	Process(data interface{}) interface{}
}

type Pipeline []Pipe
