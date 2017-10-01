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
