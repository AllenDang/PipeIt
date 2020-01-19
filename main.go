package main

import (
	"fmt"

	"github.com/AllenDang/PipeIt/pipe"
	g "github.com/AllenDang/giu"
)

var (
	input  string
	output string

	pipeline pipe.Pipeline
)

func changed() {
	output = ""

	var data interface{} = input
	for _, p := range pipeline {
		data = p.Process(data)
	}

	output = fmt.Sprint(data)
}

func buildConfigMenu(index int, configUI g.Layout) g.Layout {
	return g.Layout{
		g.ContextMenuV(fmt.Sprintf("%d-%s", index, "configMenu"), 0, configUI),
		g.ContextMenuV(fmt.Sprintf("%d-%s", index, "opMenu"), 1, g.Layout{
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
		for _, pb := range pipBuilders {
			builder := pb.Builder
			widgets = append(widgets,
				g.Selectable(pb.Name, func() {
					pipeline = append(pipeline, builder())
					changed()
				}),
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
				g.Button(fmt.Sprintf("%d-%s", i+1, p.GetName()), func() {}),
				buildConfigMenu(i, configUI),
				g.Label("->"))
		}
	}

	widgets = append(widgets, g.Button(" + ", nil), buildPipesMenu())

	return g.Line(widgets...)
}

func loop(w *g.MasterWindow) {
	g.SingleWindow(w, "pipeit", g.Layout{
		g.Label("Input - input or paste text below"),
		g.InputTextMultiline("##input", &input, -1, 200, 0, nil, changed),
		g.Group(g.Layout{
			g.Label("Pipeline"),
			buildPipeLineWidgets(pipeline),
		}),
		g.Label("Output - output text which is proceed by pipe"),
		g.InputTextMultiline("##output", &output, -1, -1, g.InputTextFlagsReadOnly, nil, nil),
	})
}

func main() {
	wnd := g.NewMasterWindow("PipeIt", 600, 400, true, nil)
	wnd.Main(loop)
}
