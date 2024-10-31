package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

type Character struct {
	rect *canvas.Rectangle
}

func NewCharacter(rect *canvas.Rectangle) *Character {
	return &Character{
		rect: rect,
	}
}

func (c *Character) HandleKeyEvent(event *fyne.KeyEvent) {
	switch event.Name {
	case fyne.KeyW:
		c.rect.Move(fyne.NewPos(c.rect.Position().X, c.rect.Position().Y-1))
	case fyne.KeyA:
		c.rect.Move(fyne.NewPos(c.rect.Position().X-1, c.rect.Position().Y))
	case fyne.KeyS:
		c.rect.Move(fyne.NewPos(c.rect.Position().X, c.rect.Position().Y+1))
	case fyne.KeyD:
		c.rect.Move(fyne.NewPos(c.rect.Position().X+1, c.rect.Position().Y))
	}
}
