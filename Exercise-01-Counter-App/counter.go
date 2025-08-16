package main

import (
	"strconv"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {

	counterWindow := app.New()

	window := counterWindow.NewWindow("Counter App")

	counter := 0
	label := widget.NewLabel("Counter: 0")

	button := widget.NewButton("Increment", func() {
		counter++
		label.SetText("Counter: " + strconv.Itoa(counter))
	})

	content := container.NewVBox(label, button)
	window.SetContent(content)
	window.Resize(fyne.NewSize(300, 200))

	window.ShowAndRun()
}

// The above code creates a simple counter application using the Fyne GUI toolkit in Go.
// It initializes a new application, creates a window with a label and a button,
// and increments the counter each time the button is clicked, updating the label accordingly.