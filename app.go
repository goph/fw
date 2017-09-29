package fw

import (
	"io"

	"github.com/go-kit/kit/log"
	"github.com/goph/emperror"
	"github.com/opentracing/opentracing-go"
)

// ApplicationOption sets options in the Application.
type ApplicationOption func(a *Application)

var defaults []ApplicationOption

// Allows to bind two or more ApplicationOption instances together.
func ApplicationOptions(opts ...ApplicationOption) ApplicationOption {
	return func(a *Application) {
		for _, opt := range opts {
			opt(a)
		}
	}
}

// Application collects all dependencies and exposes them in a single context.
type Application struct {
	logger       log.Logger
	errorHandler emperror.Handler
	tracer       opentracing.Tracer
	closers      []io.Closer
}

func NewApplication(opts ...ApplicationOption) *Application {
	app := new(Application)

	// Apply options
	for _, opt := range opts {
		opt(app)
	}

	// Apply defaults
	for _, def := range defaults {
		def(app)
	}

	return app
}
