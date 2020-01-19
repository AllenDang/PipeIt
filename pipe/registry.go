package pipe

type PipeBuilder struct {
	Name    string
	Tip     string
	Builder func() Pipe
}

var (
	pipeRegistry map[DataType][]*PipeBuilder
)

func init() {
	pipeRegistry = make(map[DataType][]*PipeBuilder)
	pipeRegistry[DataTypeString] = []*PipeBuilder{
		&PipeBuilder{"Split", "Split input string into string array", NewSplitPipe},
	}
	pipeRegistry[DataTypeStringArray] = []*PipeBuilder{
		&PipeBuilder{"Join", "Join input string array with given separator", NewJoinPipe},
		&PipeBuilder{"Match", "Match input string array with given regex", NewMatchPipe},
		&PipeBuilder{"Surround", "Add prefix and suffix to each element of input string array", NewSurroundPipe},
	}
}

func QueryPipes(byType DataType) []*PipeBuilder {
	if v, ok := pipeRegistry[byType]; ok {
		return v
	}

	return nil
}
