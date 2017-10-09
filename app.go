package fw

import (
	"io"
	"time"

	"github.com/go-kit/kit/log"
	"github.com/goph/emperror"
	"github.com/opentracing/opentracing-go"
	"go.uber.org/dig"
)

// ApplicationOption sets options in the Application.
type ApplicationOption func(a *Application)

var defaults []ApplicationOption

// Application collects all dependencies and exposes them in a single context.
type Application struct {
	container        *dig.Container
	logger           log.Logger
	errorHandler     emperror.Handler
	tracer           opentracing.Tracer
	closers          []io.Closer
	entries          map[string]interface{}
	lifecycleHooks   []Hook
	lifecycleTimeout time.Duration
}

func NewApplication(opts ...ApplicationOption) *Application {
	app := &Application{
		container: dig.New(),
		entries:   make(map[string]interface{}),
	}

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
