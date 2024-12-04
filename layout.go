package aurorium

import (
	"fyne.io/fyne/v2"
)

// TODO rename it

type CustomLayout struct {
	ExtendObject []int

	OffWidth  float32
	OffHeight float32
}

func (c *CustomLayout) Layout(objects []fyne.CanvasObject, size fyne.Size) {
	//   len(objects) - len(c.ExtendObject)

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
		fyne.Max(firstMin.Width, secondMin.Width)+c.OffWidth,
		fyne.Max(firstMin.Height, secondMin.Height)+c.OffHeight,
	)
}
