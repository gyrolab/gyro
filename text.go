/*
 *  Gyro - A modern UI toolkit for Golang
 *  Copyright (C) 2018 Roland Singer <roland@desertbit.com>
 */

package gyro

type Text interface {
	X() int
	SetX(x int)

	Y() int
	SetY(y int)

	Width() int
	Height() int

	Right() int
	Bottom() int

	Color() Color
	SetColor(c Color)

	Text() string
	SetText(t string)

	FontSize() int
	SetFontSize(s int)

	WrapLen() int
	SetWrapLen(l int)

	MouseHovered() bool

	OnMouseMove(f func(x, y int))
	OnMouseEnter(f func())
	OnMouseLeave(f func())
	OnMouseButton(f func(b MouseButton, s MouseState, x, y int))
}

func NewText() Text {
	return b.NewText()
}
