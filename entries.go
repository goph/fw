package fw

import (
	"fmt"

	"github.com/goph/emperror"
)

// Entry registers an arbitrary entry in the application.
func Entry(n string, e interface{}) Option {
	return optionFunc(func(a *Application) {
		a.entries[n] = e
	})
}

// Get returns an entry from the application.
func (a *Application) Get(name string) (interface{}, bool) {
	entry, ok := a.entries[name]

	return entry, ok
}

// MustGet returns an entry from the application and panics if it's not found.
func (a *Application) MustGet(name string) interface{} {
	entry, ok := a.Get(name)
	if !ok {
		panic(emperror.NewWithStackTrace(fmt.Sprintf("cannot find entry: %s", name)))
	}

	return entry
}
