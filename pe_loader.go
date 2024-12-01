package aurorium

import (
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/lang"
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
	//
	grid := container.NewGridWithRows(2, label1, label2)

	// label := widget.NewLabel(lang.L("pe_loader.pe_image"))

	path := widget.NewEntry()
	path.OnChanged = func(string) {
		ldr.path = path.Text
	}
	path.SetPlaceHolder(lang.L("pe_loader.pe_path.place_holder"))

	cmd := widget.NewEntry()
	cmd.OnChanged = func(string) {
		ldr.cmd = cmd.Text
	}

	// text := "https://github.com/RSSU-Shellcode/Aurorium"
	// size := ldr.ctx.app.Settings().Theme().Size(theme.SizeNameText) * 2
	//
	// ss := fyne.MeasureText(text, size, path.TextStyle)
	// ss.Height = path.CreateRenderer().MinSize().Height
	//
	// fmt.Println(ss.Height)

	// box := container.NewStack(label, path)

	form := widget.NewForm()
	form.Orientation = widget.Vertical
	form.Append("", grid)
	form.Append("", cmd)

	// grid2 := container.NewGridWithRows(2, path, cmd)
	// box := container.NewPadded(grid, form)
	ldr.box.Add(form)
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
