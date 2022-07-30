package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

type tappableIcon struct {
	widget.Icon
	OnTapped func() `json:"-"`
	Disabled bool
}

func newTappableIcon(res fyne.Resource, tapped func()) *tappableIcon {
	icon := &tappableIcon{}
	icon.ExtendBaseWidget(icon)
	icon.SetResource(res)
	icon.OnTapped = tapped
	icon.Disabled = false

	return icon
}

func (t *tappableIcon) Tapped(_ *fyne.PointEvent) {
	if t.Disabled {
		return
	}

	if t.OnTapped != nil {
		t.OnTapped()
	}
}

func (t *tappableIcon) TappedSecondary(_ *fyne.PointEvent) {
}
