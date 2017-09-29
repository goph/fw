package fw

import (
	"github.com/opentracing/opentracing-go"
)

func init() {
	defaults = append(defaults, defaultTracer)
}

// defaultTracer is an ApplicationOption that sets the default tracer.
func defaultTracer(a *Application) {
	if a.tracer == nil {
		a.tracer = opentracing.GlobalTracer()
	}
}

// Tracer returns an ApplicationOption that sets the tracer.
func Tracer(t opentracing.Tracer) ApplicationOption {
	return func(a *Application) {
		a.tracer = t
	}
}

// Tracer returns the application tracer.
func (a *Application) Tracer() opentracing.Tracer {
	return a.tracer
}
