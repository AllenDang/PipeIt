package pipe

import (
	"fmt"

	g "github.com/AllenDang/giu"
)

type SurroundPipe struct {
	prefix string
	suffix string
}

func NewSurroundPipe() Pipe {
	return &SurroundPipe{
		prefix: "'",
		suffix: "'",
	}
}

func (p *SurroundPipe) GetName() string {
	return "SR"
}

func (p *SurroundPipe) GetTip() string {
	return fmt.Sprintf("Surround each string of input string array with %s as prefix and %s as suffix", p.prefix, p.suffix)
}

func (p *SurroundPipe) GetInputType() DataType {
	return DataTypeStringArray
}

func (p *SurroundPipe) GetOutputType() DataType {
	return DataTypeStringArray
}

func (p *SurroundPipe) GetConfigUI(changed func()) g.Layout {
	return g.Layout{
		g.InputTextV("Prefix", 100, &(p.prefix), 0, nil, changed),
		g.InputTextV("Suffix", 100, &(p.suffix), 0, nil, changed),
	}
}

func (p *SurroundPipe) Process(data interface{}) interface{} {
	if strs, ok := data.([]string); ok {
		var result []string
		for _, s := range strs {
			result = append(result, fmt.Sprintf("%s%s%s", p.prefix, s, p.suffix))
		}

		return result
	}

	return []string{"Error: Surround only accepts string array as input"}
}
