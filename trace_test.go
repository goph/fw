package fw_test

import (
	"testing"

	"github.com/goph/fw"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/mocktracer"
	"github.com/stretchr/testify/assert"
)

func TestTracer(t *testing.T) {
	tracer := mocktracer.New()

	app := fw.NewApplication(fw.Tracer(tracer))

	assert.Equal(t, tracer, app.Tracer())
}

func TestApplication_Tracer(t *testing.T) {
	app := fw.NewApplication()

	assert.Equal(t, opentracing.GlobalTracer(), app.Tracer())
}
