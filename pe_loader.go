package aurorium

import (
	"fyne.io/fyne/v2/lang"
	"fyne.io/fyne/v2/widget"
)

type peLoader struct {
	pePath *widget.Entry
}

func newPELoader() *peLoader {
	pePath := widget.NewEntry()
	pePath.PlaceHolder = lang.L("pe_loader.pe_path.place_holder")

	loader := peLoader{
		pePath: pePath,
	}
	return &loader
}
