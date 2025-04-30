package component

import (
	"testing"

	"Dr.uml/backend/drawdata"
	"Dr.uml/backend/utils"
	"github.com/stretchr/testify/assert"
)

func TestGadget_GetDrawData(t *testing.T) {
	gadget, err := NewGadget(Class, utils.Point{X: 1, Y: 1})
	data, err := gadget.GetDrawData()
	assert.NoError(t, err)
	assert.NotNil(t, data)

	gdd := data.(drawdata.Gadget)
	assert.Equal(t, gdd.GadgetType, int(Class))
}
