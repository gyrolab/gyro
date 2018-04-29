/*
 *  Gyro - A modern UI toolkit for Golang
 *  Copyright (C) 2018 Roland Singer <roland@desertbit.com>
 */

package gyro

type Color struct {
	R uint8 // the red component in the range 0-255
	G uint8 // the green component in the range 0-255
	B uint8 // the blue component in the range 0-255
	A uint8 // the alpha component in the range 0-255
}

func RGBA(r, g, b, a uint8) Color {
	return Color{
		R: r,
		G: g,
		B: b,
		A: a,
	}
}
