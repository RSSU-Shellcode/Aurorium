package aurorium

import (
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func (p *Program) initLayout() {
	grid := container.NewGridWithRows(3)

	modules := container.NewGridWithColumns(3)
	modules.Add(p.peLoader.Object())
	grid.Add(modules)

	configTpl := widget.NewSelect([]string{"default"}, func(s string) {

	})
	grid.Add(configTpl)

	logger := widget.NewMultiLineEntry()
	grid.Add(logger)

	p.grid = grid
}
