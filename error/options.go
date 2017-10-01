package error

import (
	"github.com/go-kit/kit/log"
	"github.com/goph/emperror"
)

// options holds a list of options used during the error handler construction.
type options struct {
	handlers []emperror.Handler
	logger   log.Logger
}

// newOptions creates a new options instance,
// applies the provided option list and falls back to defaults where necessary.
func newOptions(opts ...HandlerOption) *options {
	o := new(options)

	for _, opt := range opts {
		opt(o)
	}

	return o
}

// HandlerOption sets options in the error handler.
type HandlerOption func(o *options)

// Logger sets a logger instance.
func Logger(l log.Logger) HandlerOption {
	return func(o *options) {
		o.logger = l
	}
}

// Handler appends a handler to the handler stack.
func Handler(h emperror.Handler) HandlerOption {
	return func(o *options) {
		o.handlers = append(o.handlers, h)
	}
}
