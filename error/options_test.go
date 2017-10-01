package error

import (
	"testing"

	"github.com/go-kit/kit/log"
	"github.com/goph/emperror"
	"github.com/stretchr/testify/assert"
)

func TestLogger(t *testing.T) {
	logger := log.NewNopLogger()
	opts := newOptions(Logger(logger))

	assert.Equal(t, logger, opts.logger)
}

func TestHandler(t *testing.T) {
	handler := new(emperror.TestHandler)
	opts := newOptions(Handler(handler))

	assert.Equal(t, handler, opts.handlers[0])
}

func TestConditional(t *testing.T) {
	handler := new(emperror.TestHandler)
	option := Handler(handler)

	t.Run("condition met", func(t *testing.T) {
		opts := newOptions(Conditional(true, option))

		assert.Equal(t, handler, opts.handlers[0])
	})

	t.Run("condition not met", func(t *testing.T) {
		opts := newOptions(Conditional(false, option))

		assert.Nil(t, opts.handlers)
	})
}
