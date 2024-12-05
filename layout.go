package aurorium

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

type HFlex struct {
	Extend  []int
	Padding float32

	OffWidth  float32
	OffHeight float32
}

func NewHFlex(extend ...int) *HFlex {
	return &HFlex{
		Extend:  extend,
		Padding: theme.Padding(),
	}
}

func (flex *HFlex) Layout(objects []fyne.CanvasObject, size fyne.Size) {
	// calculate space for extend objects
	var fixed float32
	for idx, obj := range objects {
		if !obj.Visible() {
			continue
		}
		if flex.isExtend(idx) {
			continue
		}
		fixed += obj.MinSize().Width + flex.Padding
	}
	// calculate width for extend objects
	numExtend := float32(len(flex.Extend))
	rem := size.Width - fixed
	rem -= numExtend * flex.Padding
	width := rem / numExtend
	// resize and move objects
	var x float32
	for idx, obj := range objects {
		ms := obj.MinSize()
		ms.Height = size.Height
		if flex.isExtend(idx) {
			ms.Width = width
			obj.Resize(ms)
		} else {
			obj.Resize(ms)
		}
		obj.Move(fyne.NewPos(x, 0))
		x += ms.Width + flex.Padding
	}
}

func (flex *HFlex) MinSize(objects []fyne.CanvasObject) fyne.Size {
	size := fyne.NewSize(0, 0)
	for _, obj := range objects {
		if !obj.Visible() {
			continue
		}
		ms := obj.MinSize()
		size.Width += ms.Width + flex.Padding
		size.Height = fyne.Max(size.Height, ms.Height)
	}
	size.Width -= flex.Padding
	// adjust size about offset options
	size.Width += flex.OffWidth
	size.Height += flex.OffHeight
	return size
}

func (flex *HFlex) isExtend(idx int) bool {
	for _, v := range flex.Extend {
		if v == idx {
			return true
		}
	}
	return false
}
