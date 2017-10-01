package fw

import (
	"testing"

	"github.com/go-kit/kit/log"
	"github.com/goph/emperror"
	"github.com/opentracing/opentracing-go/mocktracer"
	"github.com/stretchr/testify/assert"
)

func TestOptions(t *testing.T) {
	tracer := mocktracer.New()
	logger := log.NewNopLogger()
	handler := emperror.NewNopHandler()

	app := NewApplication(Options(
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
		app := NewApplication(Conditional(true, option))

		assert.Equal(t, logger, app.logger)
	})

	t.Run("condition not met", func(t *testing.T) {
		app := NewApplication(Conditional(false, option))

		assert.NotEqual(t, logger, app.logger)
	})
}
