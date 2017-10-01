package airbrake

import (
	"net/http"

	"github.com/airbrake/gobrake"
)

// options holds a list of options used during the error handler construction.
type options struct {
	host       string
	filters    []func(notice *gobrake.Notice) *gobrake.Notice
	httpClient *http.Client
	async      bool
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

// Host sets the target host in the notifier.
func Host(h string) HandlerOption {
	return func(o *options) {
		o.host = h
	}
}

// Filter appends a filter to the notifier.
func Filter(f func(notice *gobrake.Notice) *gobrake.Notice) HandlerOption {
	return func(o *options) {
		o.filters = append(o.filters, f)
	}
}

// HttpClient sets the HTTP Client in the notifier.
func HttpClient(c *http.Client) HandlerOption {
	return func(o *options) {
		o.httpClient = c
	}
}

// Async makes the notifier to send notices asynchronously.
func Async(a bool) HandlerOption {
	return func(o *options) {
		o.async = a
	}
}

// Conditional applies an option if the condition is true.
// This is useful to avoid using conditional logic when building the option list.
func Conditional(c bool, op HandlerOption) HandlerOption {
	return func(o *options) {
		if c {
			op(o)
		}
	}
}
