package fw

import (
	"testing"

	"github.com/go-kit/kit/log"
	"github.com/stretchr/testify/assert"
)

func TestConditional(t *testing.T) {
	logger := log.NewNopLogger()
	option := Logger(logger)

	t.Run("condition met", func(t *testing.T) {
		app := NewApplication(Conditional(true, option))

		assert.Equal(t, logger, app.logger)
	})

	t.Run("condition not met", func(t *testing.T) {
		app := NewApplication(Conditional(false, option))

		assert.NotEqual(t, logger, app.logger)
	})
}
