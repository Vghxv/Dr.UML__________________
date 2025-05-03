package components

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NewgraphMap(t *testing.T) {
	am := NewGraphMap()
	assert.NotNil(t, am)
	assert.IsType(t, &graphMap{}, am)
}
