package aurorium

import (
	"fmt"
	"net/url"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/lang"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

type peLoader struct {
	box *fyne.Container

	mode string
	arch int
	path string
	cmd  string
}

func newPELoader() *peLoader {
	loader := peLoader{
		box: container.NewVBox(),
	}
	loader.initLoadModeAndArch()
	loader.initPEPathAndCMD()
	loader.initGeneralOptions()
	return &loader
}

func (ldr *peLoader) initLoadModeAndArch() {
	labelArch := widget.NewLabel(lang.L("pe_loader.arch"))
	options := []string{"x86", "x64"}
	arch := widget.NewRadioGroup(options, func(arch string) {
		switch arch {
		case "x86":
			ldr.arch = 32
		case "x64":
			ldr.arch = 64
		default:
			panic(fmt.Sprintf("invalid architecture: %s", arch))
		}
	})
	arch.Required = true
	arch.Horizontal = true
	arch.SetSelected(options[0])

	labelLM := widget.NewLabel(lang.L("pe_loader.load_mode"))
	options = []string{"Embed", "HTTP", "File"}
	mode := widget.NewRadioGroup(options, func(mode string) {
		ldr.mode = strings.ToLower(mode)
	})
	mode.Required = true
	mode.Horizontal = true
	mode.SetSelected(options[0])

	box := container.NewHBox(labelLM, mode, labelArch, arch)
	ldr.box.Add(box)
}

func (ldr *peLoader) initPEPathAndCMD() {
	label1 := widget.NewLabel(lang.L("pe_loader.pe_image"))
	label2 := widget.NewLabel(lang.L("pe_loader.cmd_line"))
	grid1 := container.NewGridWithRows(2, label1, label2)

	open := widget.NewButton("OPEN", func() {

	})

	path := NewEntry()

	path.OnFocusGained = func() {
		open.Importance = widget.HighImportance
		open.Refresh()
	}
	path.OnFocusLost = func() {
		open.Importance = widget.MediumImportance
		open.Refresh()
	}
	path.OnChanged = func(string) {
		ldr.path = path.Text

	}
	path.SetPlaceHolder(lang.L("pe_loader.pe_path.place_holder"))
	path.Validator = func(s string) error {
		if s == "" {
			return nil
		}
		_, err := url.ParseRequestURI(s)
		return err
	}

	open.OnTapped = func() {

	}

	hBox := container.NewHBox(layout.NewSpacer(), container.NewPadded(open))
	com1 := container.NewStack(path, hBox)

	cmd := widget.NewEntry()
	cmd.OnChanged = func(string) {
		ldr.cmd = cmd.Text
	}
	grid2 := container.NewGridWithRows(2, com1, cmd)

	fmt.Println(grid1.MinSize().Height)
	fmt.Println(grid2.MinSize().Height)

	fmt.Println(path.MinSize().Height)
	fmt.Println(com1.MinSize().Height)
	fmt.Println(cmd.MinSize().Height)

	// form := widget.NewForm()
	// form.Append(lang.L("pe_loader.pe_image"), com)
	// form.Append(lang.L("pe_loader.cmd_line"), cmd)

	content := container.New(&CustomLayout{
		OffHeight: -15,
	}, grid1, grid2)

	ldr.box.Add(content)

}

func (ldr *peLoader) initGeneralOptions() {
	wait := widget.NewCheck(lang.L("pe_loader.wait_main"), func(check bool) {

	})

	instName := widget.NewEntry()

	genInst := widget.NewButton(lang.L("pe_loader.gen_inst"), func() {

	})
	saveTo := widget.NewButton(lang.L("pe_loader.save_to"), func() {

	})

	box := container.NewHBox(wait, instName, genInst, saveTo)
	ldr.box.Add(box)
}

func (ldr *peLoader) Object() fyne.CanvasObject {
	return ldr.box
}
