package fw_test

import (
	"testing"

	"fmt"

	"github.com/goph/fw"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNew(t *testing.T) {
	t.Run("ProvidesLifecycle", func(t *testing.T) {
		found := false

		_, err := fw.New(
			fw.Invoke(func(lc fw.Lifecycle) {
				assert.NotNil(t, lc)
				found = true
			}),
		)

		require.NoError(t, err)
		assert.True(t, found)
	})

	t.Run("CircularGraphReturnsError", func(t *testing.T) {
		type A struct{}
		type B struct{}

		_, err := fw.New(
			fw.Provide(func(A) B { return B{} }),
			fw.Provide(func(B) A { return A{} }),
		)

		require.Error(t, err)
		assert.Contains(t, err.Error(), "fw_test.A ->fw_test.B ->fw_test.A")
	})
}

func TestOptions(t *testing.T) {
	t.Run("OptionsComposition", func(t *testing.T) {
		var n int
		construct := func() struct{} {
			n++
			return struct{}{}
		}
		use := func(struct{}) {
			n++
		}

		_, err := fw.New(fw.Options(fw.Provide(construct), fw.Invoke(use)))

		require.NoError(t, err)
		assert.Equal(t, 2, n)
	})
}

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
