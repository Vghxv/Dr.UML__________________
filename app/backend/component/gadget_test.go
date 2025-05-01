package component

import (
	"testing"

	"Dr.uml/backend/drawdata"
	"Dr.uml/backend/utils"
	"github.com/stretchr/testify/assert"
)

func TestGadget_GetDrawData(t *testing.T) {
	gadget, err := NewGadget(Class, utils.Point{X: 1, Y: 1})
	assert.NoError(t, err)
	assert.NotNil(t, gadget)

	data, err := gadget.GetDrawData()
	assert.NoError(t, err)
	assert.NotNil(t, data)
	gdd := data.(drawdata.Gadget)
	assert.NotNil(t, gdd)

	// check default values
	assert.Equal(t, gdd.GadgetType, int(Class))
	assert.Equal(t, gdd.X, 1)
	assert.Equal(t, gdd.Y, 1)
	assert.Equal(t, gdd.Layer, 0)
	assert.Equal(t, len(gdd.Attributes), 3)
	assert.Equal(t, gdd.Attributes[0][0].Content, "Name")
	assert.Equal(t, gdd.Attributes[1][0].Content, "Attributes")
	assert.Equal(t, gdd.Attributes[2][0].Content, "Methods")
}
