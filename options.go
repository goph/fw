package fw

// Conditional applies an option if the condition is true.
// This is useful to avoid using conditional logic when building the option list.
func Conditional(c bool, op ApplicationOption) ApplicationOption {
	return func(a *Application) {
		if c {
			op(a)
		}
	}
}
