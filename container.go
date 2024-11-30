package aurorium

import (
	"fyne.io/fyne/v2/container"
)

func (p *Program) initContainer() {
	grid := container.NewGridWithColumns(3)

	grid.Add(p.peLoader.pePath)

	p.grid = grid
}
