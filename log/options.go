package log

import (
	"io"
	"os"

	"fmt"
	"github.com/go-kit/kit/log/level"
	"github.com/goph/emperror"
)

// format represents the supported log formats.
type format int

// String returns the format in a string representation.
func (f format) String() string {
	return formatMap[f]
}

const (
	JsonFormat format = iota
	LogfmtFormat
)

var (
	formatMap = map[format]string{
		JsonFormat:   "json",
		LogfmtFormat: "logfmt",
	}

	formatNameMap = map[string]format{
		"json":   JsonFormat,
		"logfmt": LogfmtFormat,
	}
)

// ParseFormat parses a string format name and returns the format or an error if the format is invalid.
func ParseFormat(formatName string) (format, error) {
	f, ok := formatNameMap[formatName]

	if !ok {
		return 0, emperror.NewWithStackTrace(fmt.Sprintf("invalid log format: %s", formatName))
	}

	return f, nil
}

var (
	// DefaultOutput is the default io.Writer where log messages are written (os.Stdout).
	DefaultOutput = os.Stdout

	// DefaultFormat is the default log format (json)
	// Available formats: json, logfmt
	DefaultFormat = JsonFormat

	// DefaultFallbackLevel is the default fallback level used in messages when a level is not defined.
	DefaultFallbackLevel level.Value = level.InfoValue()
)

// options holds a list of options used during the logger construction.
type options struct {
	output        io.Writer
	format        format
	fallbackLevel level.Value
	ctx           []interface{}
	debug         bool
}

// newOptions creates a new options instance,
// applies the provided option list and falls back to defaults where necessary.
func newOptions(opts ...LogOption) *options {
	o := new(options)

	for _, opt := range opts {
		opt(o)
	}

	// Default output
	if o.output == nil {
		o.output = DefaultOutput
	}

	// Default format
	if o.format == 0 {
		o.format = DefaultFormat
	}

	// Default fallback level
	if o.fallbackLevel == nil {
		o.fallbackLevel = DefaultFallbackLevel
	}

	return o
}

// LogOption configures the option list.
type LogOption func(o *options)

// Output sets the io.Writer output.
func Output(ou io.Writer) LogOption {
	return func(o *options) {
		o.output = ou
	}
}

// Format sets the log format in the logger.
func Format(f format) LogOption {
	return func(o *options) {
		o.format = f
	}
}

// FormatString accepts a string and tries to parse it as a log format.
// If the string cannot be parsed as a format, this method panics.
//
// Other than that it behaves like Format.
func FormatString(sf string) LogOption {
	f, err := ParseFormat(sf)
	if err != nil {
		panic(err)
	}

	return func(o *options) {
		o.format = f
	}
}

// FallbackLevel sets the default fallback level.
func FallbackLevel(l level.Value) LogOption {
	return func(o *options) {
		o.fallbackLevel = l
	}
}

// Context sets the default context.
func Context(c []interface{}) LogOption {
	return func(o *options) {
		o.ctx = c
	}
}

// With appends to the default context.
func With(kv ...interface{}) LogOption {
	return func(o *options) {
		o.ctx = append(o.ctx, kv...)
	}
}

// Debug enables/disables debug level.
func Debug(d bool) LogOption {
	return func(o *options) {
		o.debug = d
	}
}
