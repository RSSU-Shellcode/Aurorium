package aurorium

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

type Entry struct {
	widget.Entry

	OnFocusGained func()
	OnFocusLost   func()
}

func NewEntry() *Entry {
	entry := new(Entry)
	entry.Wrapping = fyne.TextWrap(fyne.TextTruncateClip)
	entry.ExtendBaseWidget(entry)
	return entry
}

func (obj *Entry) FocusGained() {
	if obj.OnFocusGained != nil {
		obj.OnFocusGained()
	}
	obj.Entry.FocusGained()
}

func (obj *Entry) FocusLost() {
	if obj.OnFocusLost != nil {
		obj.OnFocusLost()
	}
	obj.Entry.FocusLost()
}
