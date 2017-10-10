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

// optionFunc makes an Option from a function matching apply's signature.
type optionFunc func(*Application)

func (f optionFunc) apply(app *Application) { f(app) }

// Provide registers constructors in the application's dependency injection container.
//
// See the documentation at http://go.uber.org/dig for details about constructor function definitions.
func Provide(constructors ...interface{}) Option {
	return provideOption(constructors)
}

type provideOption []interface{}

func (p provideOption) apply(app *Application) {
	app.constructors = append(app.constructors, p...)
}

var defaults []Option

// Application collects all dependencies and exposes them in a single context.
type Application struct {
	container    *dig.Container
	constructors []interface{}

	logger           log.Logger
	errorHandler     emperror.Handler
	tracer           opentracing.Tracer
	closers          []io.Closer
	entries          map[string]interface{}
	lifecycleHooks   []Hook
	lifecycleTimeout time.Duration
}

func New(opts ...Option) *Application {
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
