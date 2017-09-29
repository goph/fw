package error_test

import (
	"testing"

	"bytes"
	"errors"

	"github.com/goph/fw/error"
	"github.com/goph/fw/log"
	"github.com/stretchr/testify/assert"
)

func TestHandler_Logger(t *testing.T) {
	buf := new(bytes.Buffer)
	logger := log.NewLogger(log.Output(buf), log.Format(log.LogfmtFormat))
	handler := error.NewHandler(error.Logger(logger))

	err := errors.New("error")

	handler.Handle(err)

	assert.Equal(t, "level=error msg=error\n", buf.String())
}
