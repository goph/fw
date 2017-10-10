package fw_test

import (
	"testing"

	kitlog "github.com/go-kit/kit/log"
	"github.com/goph/fw"
	"github.com/goph/fw/log"
	"github.com/stretchr/testify/assert"
)

func TestLogger(t *testing.T) {
	logger := kitlog.NewNopLogger()

	app := fw.New(fw.Logger(logger))

	assert.Equal(t, logger, app.Logger())
}

func TestApplication_Logger(t *testing.T) {
	app := fw.New()

	assert.Equal(t, log.NewLogger(), app.Logger())
}
