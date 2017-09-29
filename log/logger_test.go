package log_test

import (
	"testing"

	"bytes"

	"github.com/go-kit/kit/log/level"
	"github.com/goph/fw/log"
	"github.com/stretchr/testify/assert"
)

func TestNewLogger(t *testing.T) {
	buf := new(bytes.Buffer)
	logger := log.NewLogger(log.Output(buf))

	logger.Log()

	assert.Equal(t, "{\"level\":\"info\"}\n", buf.String())
}

func TestFormat(t *testing.T) {
	buf := new(bytes.Buffer)
	logger := log.NewLogger(log.Output(buf), log.Format(log.LogfmtFormat))

	logger.Log()

	assert.Equal(t, "level=info\n", buf.String())
}

func TestFallbackLevel(t *testing.T) {
	buf := new(bytes.Buffer)
	logger := log.NewLogger(log.Output(buf), log.FallbackLevel(level.WarnValue()))

	logger.Log()

	assert.Equal(t, "{\"level\":\"warn\"}\n", buf.String())
}

func TestContext(t *testing.T) {
	buf := new(bytes.Buffer)
	logger := log.NewLogger(log.Output(buf), log.With("key", "value"))

	logger.Log()

	assert.Equal(t, "{\"key\":\"value\",\"level\":\"info\"}\n", buf.String())
}

func TestDebug(t *testing.T) {
	t.Run("no debug", func(t *testing.T) {
		buf := new(bytes.Buffer)
		logger := log.NewLogger(log.Output(buf))

		logger.Log("level", level.DebugValue())

		assert.Equal(t, "", buf.String())
	})

	t.Run("debug", func(t *testing.T) {
		buf := new(bytes.Buffer)
		logger := log.NewLogger(log.Output(buf), log.Debug(true))

		logger.Log("level", level.DebugValue())

		assert.Equal(t, "{\"level\":\"debug\"}\n", buf.String())
	})
}
