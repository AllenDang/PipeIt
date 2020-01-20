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
		&PipeBuilder{"RegexpSplit", "Split input string into string array using regexp expression", NewRegexpSplitPipe},
		&PipeBuilder{"Fields", "Fields splits the string s around each instance of one or more consecutive white space characters", NewFieldsPipe},
	}
	pipeRegistry[DataTypeStringArray] = []*PipeBuilder{
		&PipeBuilder{"Join", "Join input string array with given separator", NewJoinPipe},
		&PipeBuilder{"Match", "Match input string array with given regex", NewMatchPipe},
		&PipeBuilder{"Surround", "Add prefix and suffix to each element of input string array", NewSurroundPipe},
		&PipeBuilder{"Replace", "Search and replace for each element of input string array", NewReplacePipe},
		&PipeBuilder{"Line", "Output input string array line by line", NewLinePipe},
		&PipeBuilder{"Trim", "Trim returns a slice of the string s with all leading and trailing Unicode code points contained in cutset removed", NewTrimPipe},
	}
}

func QueryPipes(byType DataType) []*PipeBuilder {
	if v, ok := pipeRegistry[byType]; ok {
		return v
	}

	return nil
}
