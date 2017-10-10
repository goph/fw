package fw_test

import (
	"fmt"

	"github.com/goph/fw"
)

func ExampleProvide() {
	type A struct{}
	type B struct{}
	type C struct{}

	fw.New(
		fw.Provide(func(*A, *B) *C { // Provides type *C, depends on *A and *B.
			return &C{}
		}),
	)
}

func ExampleProvide_error() {
	type A struct{}
	type B struct{}
	type C struct{}

	fw.New(
		fw.Provide(func(*A, *B) (*C, error) { // Provides type *C, depends on *A and *B, and indicates failure by returning an error.
			return &C{}, nil
		}),
	)
}

func ExampleProvide_multiple() {
	type A struct{}
	type B struct{}
	type C struct{}

	fw.New(
		fw.Provide(func(*A) (*B, *C, error) { // Provides type *B and *C, depends on *A, and can fail.
			return &B{}, &C{}, nil
		}),
	)
}

func ExampleProvide_multipleConstructors() {
	type A struct{}
	type B struct{}
	type C struct{}

	fw.New(
		fw.Provide(
			func(*B) (*C, error) { // The order of constructors does not matter.
				return &C{}, nil
			},
			func(*A) (*B, error) {
				return &B{}, nil
			},
		),
	)
}

func ExampleInvoke() {
	type A struct{}

	fw.New(
		fw.Provide(func() *A {
			return &A{}
		}),
		fw.Invoke(func(*A) {
			fmt.Print("invoked")
		}),
	)

	// Output: invoked
}
