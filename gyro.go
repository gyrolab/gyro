/*
 *  Gyro - A modern UI toolkit for Golang
 *  Copyright (C) 2018 Roland Singer <roland@desertbit.com>
 */

package gyro

var (
	b Backend
)

func RegisterBackend(backend Backend) {
	b = backend
}

// CApp returns the current application.
// There is only one application at a time.
func CApp() App {
	return b.App()
}

// Run is a shorthand for NewApp & Run.
func Run(f AppInitFunc) error {
	app, err := NewApp(f)
	if err != nil {
		return err
	}

	return app.Run()
}
