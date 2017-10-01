package fw

// Allows to bind two or more ApplicationOption instances together.
func Options(opts ...ApplicationOption) ApplicationOption {
	return func(a *Application) {
		for _, opt := range opts {
			opt(a)
		}
	}
}

// Conditional applies an option if the condition is true.
// This is useful to avoid using conditional logic when building the option list.
func Conditional(c bool, op ApplicationOption) ApplicationOption {
	return func(a *Application) {
		if c {
			op(a)
		}
	}
}
