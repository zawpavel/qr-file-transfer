package graphic

import (
	"image"
	"image/color"
	"log"

	"fyne.io/fyne/v2/canvas"
	"github.com/skip2/go-qrcode"
)

var InitialText = prepareText("this device and the mobile device must be connected to the same WiFi")
var HintText = prepareText("scan the QR code with your mobile device")
var ProcessingText = prepareText("processing, please wait...")
var ProcessingFolderText = prepareText("zipping your folder, please wait...")

func CreateQr(s string) *canvas.Image {
	if s == "" {
		emptyImg := canvas.NewImageFromImage(image.NewRGBA(image.Rect(0, 0, 256, 256)))
		emptyImg.FillMode = canvas.ImageFillOriginal
		return emptyImg
	}
	q, err := qrcode.New(s, qrcode.Medium)
	if err != nil {
		log.Fatal(err)
	}
	qr := q.Image(256)
	img := canvas.NewImageFromImage(qr)
	img.FillMode = canvas.ImageFillOriginal
	return img
}

func CreateOutputDirHint(s string) *canvas.Text {
	return prepareText("File uploaded to " + s)
}

func prepareText(text string) *canvas.Text {
	return &canvas.Text{
		Color:    color.Black,
		Text:     text,
		TextSize: 16.0,
	}
}
