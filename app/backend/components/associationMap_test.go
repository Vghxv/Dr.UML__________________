package components

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NewAssociationMap(t *testing.T) {
	am := NewAssociationMap()
	assert.NotNil(t, am)
	assert.IsType(t, &associationMap{}, am)
}
