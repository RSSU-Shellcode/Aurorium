package aurorium

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

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

type CustomLayout struct{}

func (c *CustomLayout) Layout(objects []fyne.CanvasObject, size fyne.Size) {
	if len(objects) < 2 {
		return
	}

	// First object uses its MinSize
	first := objects[0]
	firstMin := first.MinSize()
	first.Resize(firstMin)
	first.Move(fyne.NewPos(0, 0))

	// Second object takes the remaining space
	second := objects[1]
	second.Resize(fyne.NewSize(size.Width-firstMin.Width, firstMin.Height))
	second.Move(fyne.NewPos(firstMin.Width, 0))
}

func (c *CustomLayout) MinSize(objects []fyne.CanvasObject) fyne.Size {
	if len(objects) < 2 {
		return fyne.NewSize(0, 0)
	}

	// Minimum size is the sum of the first object's MinSize and the second's MinHeight
	firstMin := objects[0].MinSize()
	secondMin := objects[1].MinSize()
	return fyne.NewSize(
		fyne.Max(firstMin.Width, secondMin.Width),
		fyne.Max(firstMin.Height, secondMin.Height),
	)
}
