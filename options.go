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

// OptionFunc accepts a function which itself creates an ApplicationOption as well.
// It is useful when the inner ApplicationOption depends on the application itself (eg. requires the logger).
//
// 		app := fw.NewApplication(
//			fw.OptionFunc(func(a *fw.Application) fw.ApplicationOption {
//				logger := a.Logger()
//
//				return fw.ErrorHandler(
//					error.NewHandler(
//						error.Logger(logger),
//					),
//				)
//			}),
//		)
func OptionFunc(fn func(a *Application) ApplicationOption) ApplicationOption {
	return func(a *Application) {
		fn(a)(a)
	}
}
