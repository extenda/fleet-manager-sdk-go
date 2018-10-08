package examples

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFetchDriverPackage(t *testing.T) {
	data, err := fetchDriverPackage()

	assert.NoError(t, err)
	assert.EqualValues(t, "epson", *data.ID)
	assert.EqualValues(t, "Epson", *data.Name)
}
