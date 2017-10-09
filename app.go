package fw

import (
	"io"
	"time"

	"github.com/go-kit/kit/log"
	"github.com/goph/emperror"
	"github.com/opentracing/opentracing-go"
	"go.uber.org/dig"
)

// Option sets options in the Application.
type Option interface {
	apply(*Application)
}

type optionFunc func(*Application)

func (f optionFunc) apply(app *Application) { f(app) }

var defaults []Option

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

func NewApplication(opts ...Option) *Application {
	app := &Application{
		container: dig.New(),
		entries:   make(map[string]interface{}),
	}

	// Apply options
	for _, opt := range opts {
		opt.apply(app)
	}

	// Apply defaults
	for _, def := range defaults {
		def.apply(app)
	}

	return app
}
