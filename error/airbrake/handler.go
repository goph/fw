package airbrake

import (
	"github.com/airbrake/gobrake"
	"github.com/goph/emperror/airbrake"
)

// NewHandler returns a new Airbrake handler.
func NewHandler(projectId int64, projectKey string, opts ...HandlerOption) *airbrake.Handler {
	o := newOptions(opts...)

	notifier := gobrake.NewNotifier(projectId, projectKey)

	if o.host != "" {
		notifier.SetHost(o.host)
	}

	if o.httpClient != nil {
		notifier.Client = o.httpClient
	}

	for _, filter := range o.filters {
		notifier.AddFilter(filter)
	}

	return &airbrake.Handler{
		Notifier:          notifier,
		SendSynchronously: !o.async,
	}
}
