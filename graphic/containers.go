package graphic

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

var textContainer *fyne.Container
var qrContainer *fyne.Container

func SetContainers(text *fyne.Container, qr *fyne.Container) {
	textContainer = text
	qrContainer = qr
	UpdateText(InitialText)
	UpdateQr("")
}

func UpdateText(text *canvas.Text) {
	textContainer.RemoveAll()
	textContainer.Add(text)
	textContainer.Refresh()
}

func UpdateQr(qrString string) {
	qrContainer.RemoveAll()
	qrContainer.Add(CreateQr(qrString))
	qrContainer.Refresh()
}
