package aurorium

import (
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/lang"
	"fyne.io/fyne/v2/widget"
)

type peLoader struct {
	grid *fyne.Container

	mode string
	path string
	cmd  string
}

func newPELoader() *peLoader {
	loader := peLoader{
		grid: container.NewGridWithRows(4),
	}
	loader.initLoadMode()
	loader.initPEPath()
	loader.initCommandLine()
	return &loader
}

func (ldr *peLoader) initLoadMode() {
	label := widget.NewLabel(lang.L("pe_loader.load_mode"))

	options := []string{"Embed", "HTTP", "File"}
	modes := widget.NewRadioGroup(options, func(mode string) {
		ldr.mode = strings.ToLower(mode)
	})
	modes.Required = true
	modes.Horizontal = true
	modes.SetSelected(options[0])

	box := container.NewHBox(label, modes)
	ldr.grid.Add(box)
}

func (ldr *peLoader) initPEPath() {
	label := widget.NewLabel(lang.L("pe_loader.pe_image"))

	path := widget.NewEntry()
	path.OnChanged = func(path string) {
		ldr.path = path
	}
	path.SetPlaceHolder(lang.L("pe_loader.pe_path.place_holder"))

	grid := container.NewGridWithColumns(1, path)

	box := container.NewHBox(label, grid)
	ldr.grid.Add(box)
}

func (ldr *peLoader) initCommandLine() {
	label := widget.NewLabel(lang.L("pe_loader.cmd_line"))

	cmd := widget.NewEntry()
	cmd.OnChanged = func(cmd string) {
		ldr.cmd = cmd
	}

	box := container.NewHBox(label, cmd)
	ldr.grid.Add(box)
}

func (ldr *peLoader) Object() fyne.CanvasObject {
	return ldr.grid
}
