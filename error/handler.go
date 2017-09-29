package error

import (
	"github.com/goph/emperror"
	emperror_log "github.com/goph/emperror/log"
)

// NewHandler returns a new error handler.
func NewHandler(opts ...HandlerOption) emperror.Handler {
	o := newOptions(opts...)
	handlers := o.handlers

	if o.logger != nil {
		handlers = append(handlers, emperror_log.NewHandler(o.logger))
	}

	var handler emperror.Handler

	if len(handlers) == 0 {
		handler = emperror.NewNopHandler()
	} else if len(handlers) == 1 {
		handler = handlers[0]
	} else {
		handler = emperror.NewCompositeHandler(handlers...)
	}

	return handler
}
