package fw

import (
	"context"
	"fmt"
	"math"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-kit/kit/log/level"
)

func init() {
	defaults = append(defaults, LifecycleTimeout(defaultTimeout))
}

// defaultTimeout is used the context is created within the Application (eg. in Run).
const defaultTimeout = 15 * time.Second

// Hook is a set of lifecycleHooks callbacks, either of which can be nil.
// They are called during the application lifecycleHooks.
type Hook struct {
	PreStart  func() error
	OnStart   func(ctx context.Context, done chan<- interface{}) error
	PostStart func() error

	PreShutdown  func() error
	OnShutdown   func(ctx context.Context) error
	PostShutdown func() error
}

// SignalHook stops the application based on os signals.
var SignalHook = Hook{
	OnStart: func(ctx context.Context, done chan<- interface{}) error {
		ch := make(chan os.Signal, 1)
		signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)

		go func() {
			done <- <-ch
		}()

		return nil
	},
}

// LifecycleHook registers a lifecycle hook in the application.
func LifecycleHook(h Hook) ApplicationOption {
	return func(a *Application) {
		a.lifecycleHooks = append(a.lifecycleHooks, h)
	}
}

// LifecycleTimeout sets the default lifecycle timeout for the application.
func LifecycleTimeout(d time.Duration) ApplicationOption {
	return func(a *Application) {
		a.lifecycleTimeout = d
	}
}

// Start runs all PreStart, OnStart and PostStart hooks,
// returning immediately if it encounters an error.
func (a *Application) Start(ctx context.Context) (<-chan interface{}, error) {
	if timeout, ok := ctx.Deadline(); ok {
		level.Debug(a.logger).Log(
			"msg", "starting up with timeout",
			"timeout", math.Floor(timeout.Sub(time.Now()).Seconds()),
		)
	}

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
	if timeout, ok := ctx.Deadline(); ok {
		level.Debug(a.logger).Log(
			"msg", "shutting down with timeout",
			"timeout", math.Floor(timeout.Sub(time.Now()).Seconds()),
		)
	}

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

// Run starts the application, blocks on the signals channel, and then
// gracefully shuts the application down. It uses DefaultTimeout for the start
// and stop timeouts.
//
// See Start and Stop for application lifecycle details.
func (a *Application) Run() {
	startCtx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	done, err := a.Start(startCtx)
	if err != nil {
		a.errorHandler.Handle(err)
		return
	}

	r := <-done

	// The application stopped because of an error
	if err, ok := r.(error); ok || err != nil {
		a.errorHandler.Handle(err)
	} else if signal, ok := r.(os.Signal); ok { // The application stopped because of an os signal
		level.Info(a.logger).Log("msg", fmt.Sprintf("captured %v signal", signal))
	}

	shutdownCtx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	if err := a.Shutdown(shutdownCtx); err != nil {
		a.errorHandler.Handle(err)
	}
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
