/*
 *  Gyro - A modern UI toolkit for Golang
 *  Copyright (C) 2018 Roland Singer <roland@desertbit.com>
 */

package gyro

type Window interface {
	Close() error
	CloseChan() <-chan struct{}
	IsClosed() bool

	AddWidget(w Widget)

	// TODO
	/*
		X() int
		SetX(x int)

		Y() int
		SetY(y int)

		Position() (x int, y int)
		SetPosition(x int, y int)

		Width() int
		SetWidth(w int)

		Height() int
		SetHeight(h int)

		Size() (w int, h int)
		SetSize(w int, h int)

		Color() Color
		SetColor(c Color)

		Show()
		Hide()

		OnMouseMove(f func(x, y int))
		OnMouseButton(f func(b MouseButton, s MouseState, x, y int))*/
}

// NewWindow create a window.
func NewWindow(title string, width, height int) (Window, error) {
	return b.NewWindow(title, width, height)
}
