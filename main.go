package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/Bhargav16exd/dock-codebase.git/internal"
)

func main() {

	a := app.New()
	w := a.NewWindow("Kiosk")

	// Kiosk mode
	w.SetFullScreen(true)
	w.SetPadded(false)

	statusLabel := widget.NewLabel("Booting...")
	statusLabel.Alignment = fyne.TextAlign(fyne.TextAlignCenter)

	root := container.NewCenter(statusLabel)
	w.SetContent(root)
	w.Show()

	go startApp(statusLabel)

	a.Run()

}

func startApp(label *widget.Label) {

	update := func(text string) {
		fyne.Do(func() {
			label.SetText(text)
		})
	}

	update("Checking Networking Connection")

	for {
		if internal.IsNetworkAvailable() {
			update("Network Connected!")
			break
		}
		update("No Internet found")
	}

	//get configs
	configs := internal.GetConfig()

	if !configs.IsDockActive {
		internal.ActivateDock()
		internal.GenerateCryptoKeys()
	}

	//TBD
}
