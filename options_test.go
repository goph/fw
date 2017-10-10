package fw

import (
	"testing"

	"github.com/opentracing/opentracing-go/mocktracer"
	"github.com/stretchr/testify/assert"
)

func TestOptions(t *testing.T) {
	tracer := mocktracer.New()

	app := New(Options(
		Tracer(tracer),
	))

	assert.Equal(t, tracer, app.Tracer())
}

func TestConditional(t *testing.T) {
	tracer := mocktracer.New()
	option := Tracer(tracer)

	t.Run("condition met", func(t *testing.T) {
		app := New(Conditional(true, option))

		assert.Equal(t, tracer, app.tracer)
	})

	t.Run("condition not met", func(t *testing.T) {
		app := New(Conditional(false, option))

		assert.NotEqual(t, tracer, app.tracer)
	})
}
