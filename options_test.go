package fw

import (
	"testing"

	"github.com/go-kit/kit/log"
	"github.com/goph/emperror"
	"github.com/goph/fw/error"
	"github.com/opentracing/opentracing-go/mocktracer"
	"github.com/stretchr/testify/assert"
)

func TestOptions(t *testing.T) {
	tracer := mocktracer.New()
	logger := log.NewNopLogger()
	handler := emperror.NewNopHandler()

	app := New(Options(
		Tracer(tracer),
		Logger(logger),
		ErrorHandler(handler),
	))

	assert.Equal(t, tracer, app.Tracer())
	assert.Equal(t, logger, app.Logger())
	assert.Equal(t, handler, app.ErrorHandler())
}

func TestConditional(t *testing.T) {
	logger := log.NewNopLogger()
	option := Logger(logger)

	t.Run("condition met", func(t *testing.T) {
		app := New(Conditional(true, option))

		assert.Equal(t, logger, app.logger)
	})

	t.Run("condition not met", func(t *testing.T) {
		app := New(Conditional(false, option))

		assert.NotEqual(t, logger, app.logger)
	})
}

func TestOptionFunc(t *testing.T) {
	app := New(
		optionFunc(DefaultLogger),
		OptionFunc(func(a *Application) Option {
			logger := a.Logger()

			return ErrorHandler(
				error.NewHandler(
					error.Logger(logger),
				),
			)
		}),
	)

	assert.Equal(t, error.NewHandler(error.Logger(app.Logger())), app.ErrorHandler())
}
