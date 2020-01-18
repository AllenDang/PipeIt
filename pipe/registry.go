package pipe

type PipeBuilder struct {
	Name    string
	Builder func() Pipe
}

var (
	pipeRegistry map[DataType][]*PipeBuilder
)

func init() {
	pipeRegistry = make(map[DataType][]*PipeBuilder)
	pipeRegistry[DataTypeString] = []*PipeBuilder{
		&PipeBuilder{"Split", NewSplitPipe},
	}
	pipeRegistry[DataTypeStringArray] = []*PipeBuilder{
		&PipeBuilder{"Join", NewJoinPipe},
		&PipeBuilder{"Match", NewMatchPipe},
	}
}

func QueryPipes(byType DataType) []*PipeBuilder {
	if v, ok := pipeRegistry[byType]; ok {
		return v
	}

	return nil
}
