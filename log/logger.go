package log

import (
	"fmt"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
)

// NewLogger returns a new Logger.
func NewLogger(opts ...LogOption) log.Logger {
	o := newOptions(opts...)

	var logger log.Logger

	w := log.NewSyncWriter(o.output)

	switch o.format {
	case LogfmtFormat:
		logger = log.NewLogfmtLogger(w)

	case JsonFormat:
		logger = log.NewJSONLogger(w)

	default:
		panic(fmt.Sprintf("unsupported log format: %s", o.format))
	}

	// Add default context
	if len(o.ctx) > 0 {
		logger = log.With(logger, o.ctx...)
	}

	// Fallback to Info level
	logger = level.NewInjector(logger, o.fallbackLevel)

	// Only log debug level messages if debug mode is turned on
	if o.debug == false {
		logger = level.NewFilter(logger, level.AllowInfo())
	}

	return logger
}
