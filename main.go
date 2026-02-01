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

	//function will repeated look for files
	go internal.CheckForDataFromServer()
	go internal.CheckForFilesAvailable()

	// for {

	// 	fmt.Println("hi")

	// 	var frame internal.Frame
	// 	err := json.NewDecoder(conn).Decode(&frame)

	// 	if err != nil {
	// 		fmt.Println(err)
	// 	}

	// 	if frame.FrameMessageType == internal.MessageTypeFileMetaData.String() {
	// 		err := json.NewEncoder(conn).Encode(internal.Frame{
	// 			FrameMessageType: internal.MessageTypeAck.String(),
	// 		})
	// 		fmt.Println(err)
	// 	}

	// 	fmt.Println(frame.FileMetaData.FileName)

	// 	if frame.FrameMessageType == internal.MessageTypeFile.String() {

	// 		fmt.Println(frame.FileMetaData.FileName)

	// 		f, err := os.OpenFile(
	// 			"./backups/"+frame.FileMetaData.FileName,
	// 			os.O_CREATE|os.O_WRONLY|os.O_APPEND,
	// 			0777,
	// 		)

	// 		if err != nil {
	// 			fmt.Println(err)
	// 			return
	// 		}
	// 		defer f.Close()

	// 		_, err = f.Write(frame.Payload)

	// 	}

	// 	fmt.Println("someerr")
	// }
	// 	//get file info from connection
	// 	//if file exists
	// 	//write file on client
	// 	// on success return ack
	// 	// on fail return fail

	// 	//if no file exists , wait and restart again after some time

	//TBD
}
