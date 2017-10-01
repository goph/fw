package fw

import (
	"github.com/goph/emperror"
	"github.com/goph/fw/error"
)

func init() {
	defaults = append(defaults, DefaultErrorHandler)
}

// DefaultErrorHandler is an ApplicationOption that sets the default error handler.
func DefaultErrorHandler(a *Application) {
	if a.errorHandler == nil {
		a.errorHandler = error.NewHandler(
			error.Logger(a.Logger()),
		)
	}
}

// ErrorHandler returns an ApplicationOption that sets the error handler.
func ErrorHandler(h emperror.Handler) ApplicationOption {
	return func(a *Application) {
		a.errorHandler = h
	}
}

// ErrorHandler returns the application error handler.
func (a *Application) ErrorHandler() emperror.Handler {
	return a.errorHandler
}
