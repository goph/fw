package airbrake

import (
	"testing"

	"net/http"

	"github.com/airbrake/gobrake"
	"github.com/stretchr/testify/assert"
)

func TestHost(t *testing.T) {
	opts := newOptions(Host("host"))

	assert.Equal(t, "host", opts.host)
}

func TestFilter(t *testing.T) {
	filter := func(notice *gobrake.Notice) *gobrake.Notice { return nil }
	opts := newOptions(Filter(filter))

	assert.Equal(t, 1, len(opts.filters))
}

func TestHttpClient(t *testing.T) {
	opts := newOptions(HttpClient(http.DefaultClient))

	assert.Equal(t, http.DefaultClient, opts.httpClient)
}

func TestAsync(t *testing.T) {
	opts := newOptions(Async(true))

	assert.Equal(t, true, opts.async)
}

func TestConditional(t *testing.T) {
	t.Run("condition met", func(t *testing.T) {
		opts := newOptions(Conditional(true, Async(true)))

		assert.Equal(t, true, opts.async)
	})

	t.Run("condition not met", func(t *testing.T) {
		opts := newOptions(Conditional(false, Async(true)))

		assert.Equal(t, false, opts.async)
	})
}
