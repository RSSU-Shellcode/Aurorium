package aurorium

import (
	"fyne.io/fyne/v2/container"
)

func (p *Program) initLayout() {
	grid := container.NewGridWithRows(3)

	modules := container.NewGridWithColumns(3)
	modules.Add(p.peLoader.Object())
	grid.Add(modules)

	p.grid = grid
}
