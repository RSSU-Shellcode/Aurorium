package aurorium

import (
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/lang"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

type peLoader struct {
	ctx *Program

	box *fyne.Container

	mode string
	path string
	cmd  string
}

func newPELoader(ctx *Program) *peLoader {
	loader := peLoader{
		ctx: ctx,
		box: container.NewVBox(),
	}
	loader.initLoadMode()

	loader.initPEPathAndCMD()

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
	ldr.box.Add(box)
}

func (ldr *peLoader) initPEPathAndCMD() {
	label1 := widget.NewLabel(lang.L("pe_loader.pe_image"))
	label2 := widget.NewLabel(lang.L("pe_loader.cmd_line"))
	grid1 := container.NewGridWithRows(2, label1, label2)

	path := widget.NewEntry()
	path.OnChanged = func(string) {
		ldr.path = path.Text
	}
	path.SetPlaceHolder(lang.L("pe_loader.pe_path.place_holder"))

	open := widget.NewButton("OPEN", func() {

	})
	open.Importance = widget.HighImportance

	hBox := container.NewHBox(layout.NewSpacer(), open)
	com := container.NewPadded(hBox)

	com1 := container.NewStack(path, com)

	cmd := widget.NewEntry()
	cmd.OnChanged = func(string) {
		ldr.cmd = cmd.Text
	}
	grid2 := container.NewGridWithRows(2, com1, cmd)

	// form := widget.NewForm()
	// form.Append(lang.L("pe_loader.pe_image"), com)
	// form.Append(lang.L("pe_loader.cmd_line"), cmd)

	content := container.New(&CustomLayout{}, grid1, grid2)

	ldr.box.Add(content)

	open1 := widget.NewButton("Open", func() {

	})
	ldr.box.Add(open1)
}

func (ldr *peLoader) initCommandLine() {
	label := widget.NewLabel(lang.L("pe_loader.cmd_line"))

	cmd := widget.NewEntry()
	cmd.OnChanged = func(cmd string) {
		ldr.cmd = cmd
	}

	box := container.NewHBox(label, cmd)
	ldr.box.Add(box)
}

func (ldr *peLoader) Object() fyne.CanvasObject {
	return ldr.box
}
