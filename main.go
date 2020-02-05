package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/AllenDang/PipeIt/pipe"
	g "github.com/AllenDang/giu"
)

var (
	input  string
	output string

	inputHeight float32 = 300
	delta       float32

	pipeHint string

	pipeline pipe.Pipeline
)

func changed() {
	if len(pipeline) > 0 {
		pipeHint = "Pipeline (left click pipe to config, right click to delete)"
	} else {
		pipeHint = "Pipeline (click + to add a pipe)"
	}

	output = ""

	var data interface{} = input
	for _, p := range pipeline {
		data = p.Process(data)
	}

	output = fmt.Sprint(data)
}

func buildConfigMenu(index int, configUI g.Layout) g.Layout {
	inType := pipe.DataTypeString
	outType := pipeline[index].GetInputType()

	if index > 0 {
		inType = pipeline[index-1].GetOutputType()
	}

	betweenPipes := pipe.QueryPipesBetween(inType, outType)

	var addBeforeMenuItems g.Layout
	for i, p := range betweenPipes {
		builder := p.Builder
		addBeforeMenuItems = append(addBeforeMenuItems, g.Selectable(fmt.Sprintf("%s##%d-%d", p.Name, index, i), func() {
			pipeline = append(pipeline[:index], append(pipe.Pipeline{builder()}, pipeline[index:]...)...)
			changed()
		}))
		addBeforeMenuItems = append(addBeforeMenuItems, g.Tooltip(p.Tip))
	}

	var addBeforeMenu g.Layout
	if len(addBeforeMenuItems) > 0 {
		addBeforeMenu = append(addBeforeMenu, g.Menu(fmt.Sprintf("Add before##%d", index), addBeforeMenuItems))
	} else {
		addBeforeMenu = append(addBeforeMenu, g.Menu(fmt.Sprintf("Add before##%d", index), g.Layout{g.Label("No suitable pipe")}))
	}

	return g.Layout{
		g.Custom(func() {
			if configUI != nil {
				g.ContextMenuV(fmt.Sprintf("%s##%d", "configMenu", index), 0, configUI).Build()
			}
		}),
		g.ContextMenuV(fmt.Sprintf("%s##%d", "opMenu", index), 1, g.Layout{
			addBeforeMenu,
			g.Selectable("Delete", func() {
				pipeline = append(pipeline[:index], pipeline[index+1:]...)
				changed()
			}),
		}),
	}
}

func buildPipesMenu() g.Widget {
	var widgets []g.Widget
	queryType := pipe.DataTypeString
	if len(pipeline) > 0 {
		queryType = pipeline[len(pipeline)-1].GetOutputType()
	}

	pipBuilders := pipe.QueryPipes(queryType)
	if pipBuilders == nil {
		widgets = append(widgets, g.Label("No suitable pipe"))
	} else {
		for i, pb := range pipBuilders {
			builder := pb.Builder
			widgets = append(widgets,
				g.Selectable(fmt.Sprintf("%s##%d", pb.Name, i), func() {
					pipeline = append(pipeline, builder())
					changed()
				}),
				g.Tooltip(pb.Tip),
			)
		}
	}

	return g.ContextMenuV("AvailabePipes", 0, widgets)
}

func buildPipeLineWidgets(pipes pipe.Pipeline) g.Widget {
	var widgets []g.Widget
	if len(pipes) > 0 {
		for i, p := range pipes {
			configUI := p.GetConfigUI(func() { changed() })
			widgets = append(widgets,
				g.Button(fmt.Sprintf(" %s ##%d", p.GetName(), i), func() {}),
				g.Tooltip(p.GetTip()),
				buildConfigMenu(i, configUI),
				g.Label("->"))
		}
	}

	widgets = append(widgets, g.Button(" + ", nil), buildPipesMenu())

	return g.Line(widgets...)
}

func loop() {
	inputHeight += delta

	g.SingleWindow("pipeit", g.Layout{
		g.Label("Input - input or paste text below"),
		g.InputTextMultiline("##input", &input, -1, inputHeight, 0, nil, changed),
		g.HSplitter("hsplitter", -1, 8, &delta),
		g.Label(pipeHint),
		buildPipeLineWidgets(pipeline),
		g.Dummy(0, 8),
		g.Label("Output - output text which is processed by pipeline"),
		g.InputTextMultiline("##output", &output, -1, -1, g.InputTextFlagsReadOnly, nil, nil),
	})
}

func readStdin() {
	stat, err := os.Stdin.Stat()
	if err != nil {
		return
	}

	if (stat.Mode() & os.ModeCharDevice) == 0 {
		// Data is being piped to stdin
		bytes, _ := ioutil.ReadAll(os.Stdin)
		input = string(bytes)
	}
}

func main() {
	pipeHint = "Pipeline (click + to add a pipe)"

	// Try to read from stdin if there is anything.
	readStdin()

	wnd := g.NewMasterWindow("PipeIt", 1024, 768, true, nil)
	wnd.Main(loop)
}
