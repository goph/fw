package fw

import (
	"context"
)

// Hook is a set of lifecycleHooks callbacks, either of which can be nil.
// They are called during the application lifecycleHooks.
type Hook struct {
	PreStart  func() error
	OnStart   func(context.Context, chan<- interface{}) error
	PostStart func() error

	PreShutdown  func() error
	OnShutdown   func(context.Context) error
	PostShutdown func() error
}

// LifecycleHook registers a lifecycle hook in the application.
func LifecycleHook(h Hook) ApplicationOption {
	return func(a *Application) {
		a.lifecycleHooks = append(a.lifecycleHooks, h)
	}
}

// Start runs all PreStart, OnStart and PostStart hooks,
// returning immediately if it encounters an error.
func (a *Application) Start(ctx context.Context) (<-chan interface{}, error) {
	done := make(chan interface{}, len(a.lifecycleHooks))

	for _, hook := range a.lifecycleHooks {
		err := invokeHook(hook.PreStart)
		if err != nil {
			return done, err
		}
	}

	for _, hook := range a.lifecycleHooks {
		if hook.OnStart != nil {
			err := hook.OnStart(ctx, done)
			if err != nil {
				return done, err
			}
		}
	}

	for _, hook := range a.lifecycleHooks {
		err := invokeHook(hook.PostStart)
		if err != nil {
			return done, err
		}
	}

	return done, nil
}

// Shutdown runs all PreShutdown, OnShutdown and PostShutdown hooks,
// returning immediately if it encounters an error.
func (a *Application) Shutdown(ctx context.Context) error {
	for _, hook := range a.lifecycleHooks {
		err := invokeHook(hook.PreShutdown)
		if err != nil {
			return err
		}
	}

	for _, hook := range a.lifecycleHooks {
		err := invokeHookCtx(hook.OnShutdown, ctx)
		if err != nil {
			return err
		}
	}

	for _, hook := range a.lifecycleHooks {
		err := invokeHook(hook.PostShutdown)
		if err != nil {
			return err
		}
	}

	return nil
}

// invokeHook checks if a hook is nil first.
func invokeHook(fn func() error) error {
	if fn != nil {
		return fn()
	}

	return nil
}

// invokeHookCtx checks if a hook is nil first.
func invokeHookCtx(fn func(ctx context.Context) error, ctx context.Context) error {
	if fn != nil {
		return fn(ctx)
	}

	return nil
}
