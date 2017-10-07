package fw

import "errors"

// ErrEntryNotFound is returned when an entry is not found in the application.
var ErrEntryNotFound = errors.New("entry not found")

// Provide registers an arbitrary entry in the application.
func Provide(n string, e interface{}) ApplicationOption {
	return func(a *Application) {
		a.entries[n] = e
	}
}

// Get returns an entry from the application.
func (a *Application) Get(name string) (interface{}, error) {
	entry, ok := a.entries[name]
	if !ok {
		return nil, ErrEntryNotFound
	}

	return entry, nil
}

// MustGet returns an entry from the application and panics if it's not found.
func (a *Application) MustGet(name string) interface{} {
	entry, err := a.Get(name)
	if err != nil {
		panic(err)
	}

	return entry
}
