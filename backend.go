/*
 *  Gyro - A modern UI toolkit for Golang
 *  Copyright (C) 2018 Roland Singer <roland@desertbit.com>
 */

package gyro

// Panic on methods without error return value.
type Backend interface {
	App() App
	NewApp(f AppInitFunc) (App, error)
	NewWindow(title string, width, height int) (Window, error)
	NewRect() Rect
	NewText() Text
}
