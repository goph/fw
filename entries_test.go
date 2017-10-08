package fw_test

import (
	"testing"

	"github.com/goph/fw"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestEntry(t *testing.T) {
	app := fw.NewApplication(fw.Entry("name", "entry"))

	assert.Equal(t, "entry", app.MustGet("name"))
}

func TestApplication_Get(t *testing.T) {
	app := fw.NewApplication(fw.Entry("name", "entry"))

	entry, err := app.Get("name")

	require.NoError(t, err)
	assert.Equal(t, "entry", entry)
}

func TestApplication_Get_NotFound(t *testing.T) {
	app := fw.NewApplication()

	entry, err := app.Get("name")

	require.Error(t, err)
	assert.EqualError(t, err, "entry not found")
	assert.Nil(t, entry)
}

func TestApplication_MustGet(t *testing.T) {
	app := fw.NewApplication(fw.Entry("name", "entry"))

	assert.NotPanics(t, func() {
		entry := app.MustGet("name")

		assert.Equal(t, "entry", entry)
	})
}

func TestApplication_MustGet_NotFound(t *testing.T) {
	app := fw.NewApplication()

	assert.Panics(t, func() {
		app.MustGet("name")
	})
}
