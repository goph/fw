package log

import (
	"testing"

	"os"

	"github.com/go-kit/kit/log/level"
	"github.com/stretchr/testify/assert"
)

func TestOutput(t *testing.T) {
	w := os.Stderr
	opts := newOptions(Output(w))

	assert.Equal(t, w, opts.output)
}

func TestOutput_Default(t *testing.T) {
	opts := newOptions()

	assert.Equal(t, DefaultOutput, opts.output)
}

func TestDefaultOutput(t *testing.T) {
	assert.Equal(t, os.Stdout, DefaultOutput)
}

func TestFormat(t *testing.T) {
	opts := newOptions(Format(LogfmtFormat))

	assert.Equal(t, LogfmtFormat, opts.format)
}

func TestFormat_Default(t *testing.T) {
	opts := newOptions()

	assert.Equal(t, DefaultFormat, opts.format)
}

func TestDefaultFormat(t *testing.T) {
	assert.Equal(t, JsonFormat, DefaultFormat)
}

func TestFallbackLevel(t *testing.T) {
	lvl := level.WarnValue()
	opts := newOptions(FallbackLevel(lvl))

	assert.Equal(t, lvl, opts.fallbackLevel)
}

func TestFallbackLevel_Default(t *testing.T) {
	opts := newOptions()

	assert.Equal(t, DefaultFallbackLevel, opts.fallbackLevel)
}

func TestDefaultFallbackLevel(t *testing.T) {
	assert.Equal(t, level.InfoValue(), DefaultFallbackLevel)
}

func TestContext(t *testing.T) {
	ctx := []interface{}{"key", "value"}
	opts := newOptions(Context(ctx))

	assert.Equal(t, ctx, opts.ctx)
}

func TestContext_Default(t *testing.T) {
	opts := newOptions()

	assert.Nil(t, opts.ctx)
}

func TestWith(t *testing.T) {
	ctx := []interface{}{"key", "value"}
	opts := newOptions(With(ctx...))

	assert.Equal(t, ctx, opts.ctx)
}

func TestWithContext(t *testing.T) {
	ctx := []interface{}{"key", "value"}
	ctx2 := []interface{}{"key2", "value2"}
	opts := newOptions(Context(ctx), With(ctx2...))

	assert.Equal(t, append(ctx, ctx2...), opts.ctx)
}
