/*
 *  Gyro - A modern UI toolkit for Golang
 *  Copyright (C) 2018 Roland Singer <roland@desertbit.com>
 */

package gyro

type Rect interface {
	AddWidget(w Widget)

	X() int
	SetX(x int)

	Y() int
	SetY(y int)

	Width() int
	SetWidth(w int)

	Height() int
	SetHeight(h int)

	Right() int
	Bottom() int

	Radius() int
	SetRadius(r int)

	Border() int
	SetBorder(b int)

	Color() Color
	SetColor(c Color)

	BorderColor() Color
	SetBorderColor(c Color)

	/*
		MouseHovered() bool

		OnMouseMove(f func(x, y int))
		OnMouseEnter(f func())
		OnMouseLeave(f func())
		OnMouseButton(f func(b MouseButton, s MouseState, x, y int))
	*/
}

func NewRect() Rect {
	return b.NewRect()
}
