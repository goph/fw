package fw

import (
	kitlog "github.com/go-kit/kit/log"
	"github.com/goph/fw/log"
)

func init() {
	// Prepend the default logger
	defaults = append([]Option{optionFunc(DefaultLogger)}, defaults...)
}

// DefaultLogger is an ApplicationOption that sets the default logger.
func DefaultLogger(a *Application) {
	if a.logger == nil {
		a.logger = log.NewLogger()
	}
}

// Logger returns an ApplicationOption that sets the logger.
func Logger(l kitlog.Logger) Option {
	return optionFunc(func(a *Application) {
		a.logger = l
	})
}

// Logger returns the application logger.
func (a *Application) Logger() kitlog.Logger {
	return a.logger
}
