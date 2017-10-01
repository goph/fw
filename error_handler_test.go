package fw_test

import (
	"testing"

	"github.com/go-kit/kit/log"
	"github.com/goph/emperror"
	"github.com/goph/fw"
	"github.com/goph/fw/error"
	"github.com/stretchr/testify/assert"
)

func TestErrorHandler(t *testing.T) {
	handler := emperror.NewNopHandler()

	app := fw.NewApplication(fw.ErrorHandler(handler))

	assert.Equal(t, handler, app.ErrorHandler())
}

func TestApplication_ErrorHandler(t *testing.T) {
	logger := log.NewNopLogger()

	app := fw.NewApplication(fw.Logger(logger))

	assert.Equal(t, error.NewHandler(error.Logger(logger)), app.ErrorHandler())
}
