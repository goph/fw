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

// Application collects all dependencies and exposes them in a single context.
type Application struct {
	logger       log.Logger
	errorHandler emperror.Handler
	tracer       opentracing.Tracer
	closers      []io.Closer
	entries      map[string]interface{}
}

func NewApplication(opts ...ApplicationOption) *Application {
	app := new(Application)
	app.entries = make(map[string]interface{})

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
