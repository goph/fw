package fw_test

import (
	"testing"

	"context"

	"github.com/goph/fw"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLifecycleHook(t *testing.T) {
	var preStart, onStart, postStart, preShutdown, onShutdown, postShutdown bool

	hook := fw.Hook{
		PreStart:  testHook(&preStart),
		OnStart:   func(ctx context.Context, done chan<- interface{}) error {
			onStart = true

			return nil
		},
		PostStart: testHook(&postStart),

		PreShutdown:  testHook(&preShutdown),
		OnShutdown:   testHookCtx(&onShutdown),
		PostShutdown: testHook(&postShutdown),
	}

	app := fw.NewApplication(fw.LifecycleHook(hook))

	_, startErr := app.Start(context.Background())
	shutdownErr := app.Shutdown(context.Background())

	require.NoError(t, startErr)
	require.NoError(t, shutdownErr)

	assert.True(t, preStart)
	assert.True(t, onStart)
	assert.True(t, postStart)

	assert.True(t, preShutdown)
	assert.True(t, onShutdown)
	assert.True(t, postShutdown)
}

func testHook(assertion *bool) func() error {
	return func() error {
		*assertion = true

		return nil
	}
}

func testHookCtx(assertion *bool) func(ctx context.Context) error {
	return func(ctx context.Context) error {
		*assertion = true

		return nil
	}
}
