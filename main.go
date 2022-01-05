package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/AllenDang/PipeIt/pipe"
	g "github.com/AllenDang/giu"
)

const (
	saveDir string = "Save"
)

var (
	input  string
	output string

	pipeHint string

	pipeline pipe.Pipeline

	savePipelineName string
	savedPipelines   []string
	selectedIndex    int32
	comboPreview     string
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
		addBeforeMenuItems = append(addBeforeMenuItems, g.Selectable(fmt.Sprintf("%s##%d-%d", p.Name, index, i)).OnClick(func() {
			pipeline = append(pipeline[:index], append(pipe.Pipeline{builder()}, pipeline[index:]...)...)
			changed()
		}))
		addBeforeMenuItems = append(addBeforeMenuItems, g.Tooltip(p.Tip))
	}

	var addBeforeMenu g.Layout
	if len(addBeforeMenuItems) > 0 {
		addBeforeMenu = append(addBeforeMenu, g.Menu(fmt.Sprintf("Add before##%d", index)).Layout(addBeforeMenuItems))
	} else {
		addBeforeMenu = append(addBeforeMenu, g.Menu(fmt.Sprintf("Add before##%d", index)).Layout(g.Layout{g.Label("No suitable pipe")}))
	}

	return g.Layout{
		g.Custom(func() {
			if configUI != nil {
				g.ContextMenu().ID(fmt.Sprintf("%s##%d", "configMenu", index)).MouseButton(g.MouseButtonLeft).Layout(configUI).Build()
			}
		}),
		g.ContextMenu().ID(fmt.Sprintf("%s##%d", "opMenu", index)).MouseButton(g.MouseButtonRight).Layout(g.Layout{
			addBeforeMenu,
			g.Selectable("Delete").OnClick(func() {
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
				g.Selectable(fmt.Sprintf("%s##%d", pb.Name, i)).OnClick(func() {
					pipeline = append(pipeline, builder())
					changed()
				}),
				g.Tooltip(pb.Tip),
			)
		}
	}

	return g.ContextMenu().ID("AvailabePipes").MouseButton(g.MouseButtonLeft).Layout(widgets...)
}

func buildPipeLineWidgets(pipes pipe.Pipeline) g.Widget {
	var widgets []g.Widget
	if len(pipes) > 0 {
		for i, p := range pipes {
			configUI := p.GetConfigUI(func() { changed() })
			widgets = append(widgets,
				g.Button(fmt.Sprintf(" %s ##%d", p.GetName(), i)),
				g.Tooltip(p.GetTip()),
				buildConfigMenu(i, configUI),
				g.Label("->"))
		}
	}

	widgets = append(widgets, g.Button(" + "), buildPipesMenu())

	return g.Row(widgets...)
}

func btnLoadClicked() {
	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	dir = filepath.Join(dir, saveDir)

	f, err := os.Open(filepath.Join(dir, comboPreview))
	if err != nil {
		g.Msgbox("Error", fmt.Sprintf("Load pipeline failed, error message is %s", err.Error()))
		return
	}
	defer f.Close()

	pl, err := pipe.DecodePipeline(f)
	if err != nil {
		g.Msgbox("Error", fmt.Sprintf("Load pipeline failed, error message is %s", err.Error()))
		return
	}

	pipeline = *pl
	changed()
}

func btnSaveClicked() {
	if len(pipeline) == 0 {
		g.Msgbox("Error", "Current pipeline is empty.")
		return
	}

	g.OpenPopup("Save Pipeline")
}

func onSave() {
	defer func() {
		g.CloseCurrentPopup()
		loadSavedPiplines()
	}()

	if len(savePipelineName) == 0 {
		g.Msgbox("Error", "Pipeline's name cannot be empty")
		return
	}

	// Prepare save dir
	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	dir = filepath.Join(dir, saveDir)

	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err := os.Mkdir(dir, os.ModeDir)
		if err != nil {
			g.Msgbox("Error", fmt.Sprintf("Failed to create save directory, error message is %s", err.Error()))
			return
		}
	}

	saveFilepath := filepath.Join(dir, fmt.Sprintf("%s.pl", savePipelineName))

	buf, err := pipe.EncodePipeline(pipeline)
	if err != nil {
		g.Msgbox("Error", fmt.Sprintf("Save pipeline failed, error message is %s", err.Error()))
		return
	}

	err = ioutil.WriteFile(saveFilepath, buf.Bytes(), 0644)
	if err != nil {
		g.Msgbox("Error", fmt.Sprintf("Save pipeline failed, error message is %s", err.Error()))
		return
	}
}

func onCancel() {
	g.CloseCurrentPopup()
}

func onComboChanged() {
	comboPreview = savedPipelines[selectedIndex]
}

func loadSavedPiplines() {
	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	dir = filepath.Join(dir, saveDir)

	// Get all saved *.pl filenames.
	var files []string
	_ = filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err == nil && !info.IsDir() && filepath.Ext(path) == ".pl" {
			files = append(files, filepath.Base(path))
		}
		return nil
	})

	savedPipelines = files
	if int(selectedIndex) > len(savedPipelines) {
		selectedIndex = 0
	}
	if len(savedPipelines) > 0 {
		comboPreview = savedPipelines[selectedIndex]
	}
}

func loop() {
	g.SingleWindow().Layout(g.Layout{
		g.SplitLayout(g.DirectionVertical, 300,
			g.Layout{
				g.Label("Input - input or paste text below"),
				g.InputTextMultiline(&input).Size(-1, -1).OnChange(changed),
			},
			g.Layout{
				g.Dummy(0, 8),
				g.Row(
					g.Label(pipeHint),
					g.Combo("##savedPipeines", comboPreview, savedPipelines, &selectedIndex).Size(200).OnChange(onComboChanged),
					g.Button("Load").OnClick(btnLoadClicked),
					g.Button("Save").OnClick(btnSaveClicked),
				),
				g.PopupModal("Save Pipeline").Flags(g.WindowFlagsNoResize).Layout(g.Layout{
					g.Label("Enter the name of the pipeline "),
					g.InputText(&savePipelineName).Size(200),
					g.Row(
						g.Button("Save").OnClick(onSave),
						g.Button("Cancel").OnClick(onCancel),
					),
				}),
				buildPipeLineWidgets(pipeline),
				g.Dummy(0, 8),
				g.Label("Output - output text which is processed by pipeline"),
				g.InputTextMultiline(&output).Size(-1, -1).Flags(g.InputTextFlagsReadOnly),
			}).Border(false),
		g.PrepareMsgbox(),
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

	// Load saved pipelines.
	loadSavedPiplines()

	wnd := g.NewMasterWindow("PipeIt", 1024, 768, 0)
	wnd.Run(loop)
}
