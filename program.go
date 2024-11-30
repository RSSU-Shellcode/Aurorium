package aurorium

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
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
		return err
	}

	p.peLoader = newPELoader()

	p.initContainer()

	w := p.app.NewWindow("Aurorium")
	w.SetContent(p.grid)
	w.Show()

	p.app.Run()
	return nil
}
