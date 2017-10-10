package fw

import (
	"time"

	"github.com/go-kit/kit/log"
	"github.com/goph/emperror"
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

// Logger sets the application logger used for logging application lifecycle.
func Logger(logger log.Logger) Option {
	return optionFunc(func(app *Application) {
		app.logger = logger
	})
}

// ErrorHandler sets the error handler in the application.
func ErrorHandler(handler emperror.Handler) Option {
	return optionFunc(func(a *Application) {
		a.errorHandler = handler
	})
}

// Application collects all dependencies and exposes them in a single context.
type Application struct {
	container    *dig.Container
	constructors []interface{}

	logger       log.Logger
	errorHandler emperror.Handler

	lifecycleHooks   []Hook
	lifecycleTimeout time.Duration
}

func New(opts ...Option) (*Application, error) {
	app := &Application{
		container: dig.New(),
	}

	// Apply options
	for _, opt := range opts {
		opt.apply(app)
	}

	// Default logger
	if app.logger == nil {
		app.logger = log.NewNopLogger()
	}

	// Default error handler
	if app.errorHandler == nil {
		app.errorHandler = emperror.NewNopHandler()
	}

	// Default lifecycle timeout
	if app.lifecycleTimeout == 0 {
		app.lifecycleTimeout = defaultTimeout
	}

	// Register the constructors in the container
	for _, ctor := range app.constructors {
		err := app.container.Provide(ctor)
		if err != nil {
			err = emperror.WithStack(emperror.WithMessage(err, "failed to register constructor in the container"))

			return nil, err
		}
	}

	return app, nil
}
