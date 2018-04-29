/*
 *  Gyro - A modern UI toolkit for Golang
 *  Copyright (C) 2018 Roland Singer <roland@desertbit.com>
 */

package gyro

type Widget interface {
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
}
