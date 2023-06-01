package main

import (
	"log"
	"net/url"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/zawpavel/qr-file-transfer/filetransfer"
	"github.com/zawpavel/qr-file-transfer/graphic"
)

func main() {
	myApp := app.New()
	myApp.Settings().SetTheme(theme.LightTheme())
	myWindow := myApp.NewWindow("QR file transfer")

	spacer := layout.NewSpacer()

	textContainer := container.NewCenter()
	qrContainer := container.NewCenter()
	graphic.SetContainers(textContainer, qrContainer)

	container := container.NewVBox(
		spacer,
		createSendFileButton(&myWindow),
		spacer,
		createSendFolderButton(&myWindow),
		spacer,
		createReceiveButton(&myWindow),
		spacer,
		textContainer,
		spacer,
		qrContainer,
		spacer,
		createDonateButton(&myWindow, myApp),
		spacer,
		spacer,
	)

	myWindow.SetContent(container)
	myWindow.Resize(fyne.NewSize(600, 600))
	myWindow.ShowAndRun()
}

func createSendFileButton(window *fyne.Window) *widget.Button {
	button := widget.NewButton("send file", func() {
		graphic.UpdateQr("")
		graphic.UpdateText(graphic.ProcessingText)
		dialog.ShowFileOpen(func(file fyne.URIReadCloser, err error) {
			if err != nil {
				log.Println("Error:", err)
				return
			}
			if file == nil {
				graphic.UpdateText(graphic.InitialText)
				log.Println("No file selected")
				return
			}
			filepath := file.URI().Path()
			downloadLink := filetransfer.GetDownloadLink(filepath)
			graphic.UpdateQr(downloadLink)
			graphic.UpdateText(graphic.HintText)
		}, *window)
	})
	button.Importance = widget.HighImportance
	return button
}

func createSendFolderButton(window *fyne.Window) *widget.Button {
	button := widget.NewButton("send directory", func() {
		graphic.UpdateQr("")
		graphic.UpdateText(graphic.ProcessingFolderText)
		dialog.ShowFolderOpen(func(folder fyne.ListableURI, err error) {
			if err != nil {
				log.Println("Error:", err)
				return
			}
			if folder == nil {
				graphic.UpdateText(graphic.InitialText)
				log.Println("No folder selected")
				return
			}
			folderpath := folder.Path()
			log.Println("Selected folder:", folderpath)
			downloadLink := filetransfer.GetDownloadLink(folderpath)
			graphic.UpdateQr(downloadLink)
			graphic.UpdateText(graphic.HintText)
		}, *window)
	})
	button.Importance = widget.HighImportance
	return button
}

func createReceiveButton(window *fyne.Window) *widget.Button {
	button := widget.NewButton("receive", func() {
		receiveUrlChannel := make(chan string)
		outputDirChannel := make(chan string)
		go filetransfer.ReceiveFiles(receiveUrlChannel, outputDirChannel)
		receiveUrl := <-receiveUrlChannel
		graphic.UpdateText(graphic.HintText)
		graphic.UpdateQr(receiveUrl)

		// wait until file is received
		go func() {
			outDir := <-outputDirChannel
			graphic.UpdateText(graphic.CreateOutputDirHint(outDir))
			graphic.UpdateQr("")
		}()
	})
	button.Importance = widget.HighImportance
	return button
}

func createDonateButton(window *fyne.Window, app fyne.App) *widget.Button {
	button := widget.NewButton("donate", func() {
		path, _ := url.Parse("https://www.buymeacoffee.com/zawpavel")
		app.OpenURL(path)
	})
	button.Importance = widget.HighImportance
	return button
}
