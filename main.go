package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

/**
 * DateHack: A 2D Hack & Slash / Dating Simulator cross-over game written in Go using Fyne.
 *
 * @author: Scalar Studios
 */
func main() {
	// Create a new application
	a := app.New()
	w := a.NewWindow("DateHack")

	// Load a character (as a square for now) and add it to the window - add keyboard controls for movement
	rect := canvas.NewRectangle(color.White)
	rect.SetMinSize(fyne.NewSize(32, 32))

	// Center the rectangle in the window
	centeredRect := container.NewCenter(rect)
	w.SetContent(centeredRect)

	// Set the window size and disable resizing
	w.Resize(fyne.NewSize(1280, 720))
	w.SetFixedSize(true)

	// Add keyboard controls for movement
	w.Canvas().SetOnTypedKey(func(event *fyne.KeyEvent) {
		switch event.Name {
		case fyne.KeyW:
			rect.Move(fyne.NewPos(rect.Position().X, rect.Position().Y-1))
		case fyne.KeyA:
			rect.Move(fyne.NewPos(rect.Position().X-1, rect.Position().Y))
		case fyne.KeyS:
			rect.Move(fyne.NewPos(rect.Position().X, rect.Position().Y+1))
		case fyne.KeyD:
			rect.Move(fyne.NewPos(rect.Position().X+1, rect.Position().Y))
		}
	})

	// Show the window and run the application
	w.ShowAndRun()
}
