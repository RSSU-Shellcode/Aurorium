package aurorium

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// Program contains all context data.
type Program struct {
	app fyne.App

	grid *fyne.Container

	peLoader *peLoader
}

// New is used to create a new GUI program.
func New() *Program {
	p := Program{
		app: app.NewWithID("aurorium"),
	}
	return &p
}

// Main is used to run the program.
func (p *Program) Main() error {
	err := p.initLanguage()
	if err != nil {
		return fmt.Errorf("failed to inititalize language: %s", err)
	}
	err = p.initTheme()
	if err != nil {
		return fmt.Errorf("failed to inititalize theme: %s", err)
	}

	p.peLoader = newPELoader()
	p.initLayout()

	w := p.app.NewWindow("Aurorium")
	w.Resize(fyne.Size{
		Width:  1280,
		Height: 720,
	})
	w.CenterOnScreen()
	w.SetContent(p.grid)
	w.Show()

	p.app.Run()
	return nil
}

func (p *Program) initLayout() {
	grid := container.NewGridWithRows(3)

	modules := container.NewGridWithColumns(2)
	modules.Add(p.peLoader.Object())
	grid.Add(modules)

	configTpl := widget.NewSelect([]string{"default"}, func(s string) {

	})
	grid.Add(configTpl)

	logger := widget.NewMultiLineEntry()
	grid.Add(logger)

	p.grid = grid
}
