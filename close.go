package fw

import (
	"io"

	"github.com/goph/emperror"
)

// Closer returns an ApplicationOption that appends a closer to the Application's closer list.
func Closer(c io.Closer) Option {
	return optionFunc(func(a *Application) {
		a.closers = append(a.closers, c)
	})
}

// Close implements the common closer interface and closes the underlying resources.
// The resources are closed in a reversed order (just like how subsequent defer Close() calls would work).
// Errors are suppressed (again, like in case of defer calls).
func (a *Application) Close() error {
	err := emperror.Recover(recover())
	if err != nil {
		a.errorHandler.Handle(err)
	}

	// TODO: log application closing and handle errors?
	if len(a.closers) == 0 {
		return nil
	}

	// Closing resources in a reversed order
	for i := len(a.closers) - 1; i >= 0; i-- {
		a.closers[i].Close()
	}

	return nil
}
