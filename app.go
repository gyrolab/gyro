/*
 *  Gyro - A modern UI toolkit for Golang
 *  Copyright (C) 2018 Roland Singer <roland@desertbit.com>
 */

package gyro

type AppInitFunc func() error

type App interface {
	Run() error

	Close() error
	CloseChan() <-chan struct{}
	IsClosed() bool

	Sync(f func())
	Lock()
	Unlock()
}

func NewApp(f AppInitFunc) (App, error) {
	return b.NewApp(f)
}
