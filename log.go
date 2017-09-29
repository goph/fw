package fw

import (
	kitlog "github.com/go-kit/kit/log"
	"github.com/goph/fw/log"
)

func init() {
	// Prepend the default logger
	defaults = append([]ApplicationOption{defaultLogger}, defaults...)
}

// defaultLogger is an ApplicationOption that sets the default logger.
func defaultLogger(a *Application) {
	if a.logger == nil {
		a.logger = log.NewLogger()
	}
}

// Logger returns an ApplicationOption that sets the logger.
func Logger(l kitlog.Logger) ApplicationOption {
	return func(a *Application) {
		a.logger = l
	}
}

// Logger returns the application logger.
func (a *Application) Logger() kitlog.Logger {
	return a.logger
}
