package pipe

import (
	"encoding/gob"
	"fmt"
	"strings"

	g "github.com/AllenDang/giu"
)

type SurroundPipe struct {
	Prefix string
	Suffix string
}

func init() {
	gob.Register(&SurroundPipe{})
}

func NewSurroundPipe() Pipe {
	return &SurroundPipe{
		Prefix: "'",
		Suffix: "'",
	}
}

func (p *SurroundPipe) GetName() string {
	return "SR"
}

func (p *SurroundPipe) GetTip() string {
	return fmt.Sprintf("Surround each string of input string array with %s as Prefix and %s as Suffix", p.Prefix, p.Suffix)
}

func (p *SurroundPipe) GetInputType() DataType {
	return DataTypeStringArray
}

func (p *SurroundPipe) GetOutputType() DataType {
	return DataTypeStringArray
}

func (p *SurroundPipe) GetConfigUI(changed func()) g.Layout {
	return g.Layout{
		g.Label("Use %d to generate series number"),
		g.InputText(&(p.Prefix)).Label("Prefix").Size(100).OnChange(changed),
		g.InputText(&(p.Suffix)).Label("Suffix").Size(100).OnChange(changed),
	}
}

func (p *SurroundPipe) Process(data interface{}) interface{} {
	if strs, ok := data.([]string); ok {
		var result []string
		for i, s := range strs {
			pf := strings.Replace(p.Prefix, "%d", fmt.Sprintf("%d", i+1), -1)
			sf := strings.Replace(p.Suffix, "%d", fmt.Sprintf("%d", i+1), -1)
			result = append(result, fmt.Sprintf("%s%s%s", pf, s, sf))
		}

		return result
	}

	return []string{"Error: Surround only accepts string array as input"}
}
