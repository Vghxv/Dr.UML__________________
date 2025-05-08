package component

import (
	"strconv"
	"testing"

	"Dr.uml/backend/drawdata"
	"Dr.uml/backend/utils"
	"Dr.uml/backend/utils/duerror"
	"github.com/stretchr/testify/assert"
)

// test util
func newEmptyGadget(gadgetType GadgetType, point utils.Point) *Gadget {
	g, err := NewGadget(gadgetType, point)
	if err != nil {
		panic(err)
	}
	for i, length := range g.GetAttributesLen() {
		for j := 0; j < length; j++ {
			if err := g.RemoveAttribute(i, 0); err != nil {
				panic(err)
			}
		}
	}
	return g
}

type mockParent struct {
	Times int
}

func (m *mockParent) UpdateParentDraw() duerror.DUError {
	m.Times++
	return nil
}

// Constructor
func TestNewGadget(t *testing.T) {
	// success
	g, err := NewGadget(Class, utils.Point{X: 1, Y: 1})
	assert.NoError(t, err)
	assert.NotNil(t, g)
	assert.Equal(t, Class, g.GetGadgetType())

	// invalid gadget type
	g, err = NewGadget(-1, utils.Point{X: 1, Y: 1})
	assert.Error(t, err)
}

// Getter
func TestGetPoint(t *testing.T) {
	g := newEmptyGadget(Class, utils.Point{X: 1, Y: 1})
	assert.Equal(t, utils.Point{X: 1, Y: 1}, g.GetPoint())
}

func TestGetLayer(t *testing.T) {
	g := newEmptyGadget(Class, utils.Point{X: 1, Y: 1})
	assert.Equal(t, 0, g.GetLayer())
}

func TestGetColor(t *testing.T) {
	g := newEmptyGadget(Class, utils.Point{X: 1, Y: 1})
	assert.Equal(t, utils.FromHex(drawdata.DefaultGadgetColor), g.GetColor())
}

func TestGetGadgetType(t *testing.T) {
	g := newEmptyGadget(Class, utils.Point{X: 1, Y: 1})
	assert.Equal(t, Class, g.GetGadgetType())
}

func TestGetAttributesLen(t *testing.T) {
	// for Class type
	g := newEmptyGadget(Class, utils.Point{X: 1, Y: 1})
	assert.Equal(t, []int{0, 0, 0}, g.GetAttributesLen())
}

// Setter
func TestSetPoint(t *testing.T) {
	g := newEmptyGadget(Class, utils.Point{X: 1, Y: 1})
	assert.NoError(t, g.SetPoint(utils.Point{X: 2, Y: 2}))
	assert.Equal(t, utils.Point{X: 2, Y: 2}, g.GetPoint())

	mp := mockParent{}
	g.RegisterUpdateParentDraw(mp.UpdateParentDraw)
	assert.NoError(t, g.SetPoint(utils.Point{X: 3, Y: 3}))
	assert.Equal(t, utils.Point{X: 3, Y: 3}, g.GetPoint())
	assert.Equal(t, 1, mp.Times)
}

func TestSetLayer(t *testing.T) {
	g := newEmptyGadget(Class, utils.Point{X: 1, Y: 1})
	assert.NoError(t, g.SetLayer(1))
	assert.Equal(t, 1, g.GetLayer())

	mp := mockParent{}
	g.RegisterUpdateParentDraw(mp.UpdateParentDraw)
	assert.NoError(t, g.SetLayer(2))
	assert.Equal(t, 2, g.GetLayer())
	assert.Equal(t, 1, mp.Times)
}

func TestSetColor(t *testing.T) {
	g := newEmptyGadget(Class, utils.Point{X: 1, Y: 1})
	assert.NoError(t, g.SetColor(utils.FromHex(0xFF0000)))
	assert.Equal(t, utils.FromHex(0xFF0000), g.GetColor())

	mp := mockParent{}
	g.RegisterUpdateParentDraw(mp.UpdateParentDraw)
	assert.NoError(t, g.SetColor(utils.FromHex(0x00FF00)))
	assert.Equal(t, utils.FromHex(0x00FF00), g.GetColor())
	assert.Equal(t, 1, mp.Times)
}

// Methods
func TestCover(t *testing.T) {
	g, _ := NewGadget(Class, utils.Point{X: 1, Y: 1})
	width := g.GetDrawData().(drawdata.Gadget).Width
	height := g.GetDrawData().(drawdata.Gadget).Height

	// coners
	coners := []utils.Point{
		{X: 1, Y: 1},                  // top-left
		{X: 1 + width, Y: 1},          // top-right
		{X: 1, Y: 1 + height},         // bottom-left
		{X: 1 + width, Y: 1 + height}, // bottom-right
	}
	for _, corner := range coners {
		val, err := g.Cover(corner)
		assert.NoError(t, err)
		assert.True(t, val)
	}

	// inside
	val, err := g.Cover(utils.Point{X: 1 + width/2, Y: 1 + height/2})
	assert.NoError(t, err)
	assert.True(t, val)

	// outside
	outsides := []utils.Point{
		{X: 1, Y: 0},
		{X: 1, Y: 1 + height + 1},
		{X: 0, Y: 1},
		{X: 1 + width + 1, Y: 1},
	}
	for _, outside := range outsides {
		val, err := g.Cover(outside)
		assert.NoError(t, err)
		assert.False(t, val)
	}
}

func TestAddAttribute(t *testing.T) {
	g := newEmptyGadget(Class, utils.Point{X: 1, Y: 1})

	// add attribute in every section
	sectionLen := len(g.GetAttributesLen())
	for i := 0; i < sectionLen; i++ {
		content := "test" + strconv.Itoa(i)
		assert.NoError(t, g.AddAttribute(i, content))
		assert.Equal(t, 1, g.GetAttributesLen()[i])
		att_content := g.GetDrawData().(drawdata.Gadget).Attributes[i][0].Content
		assert.Equal(t, content, att_content)
	}

	// invalid section
	assert.Error(t, g.AddAttribute(-1, "test"))
	assert.Error(t, g.AddAttribute(sectionLen, "test"))

	// some errors are hard to test :(
}

func TestRemoveAttribute(t *testing.T) {
	g := newEmptyGadget(Class, utils.Point{X: 1, Y: 1})

	// every section
	sectionLen := len(g.GetAttributesLen())
	for i := 0; i < sectionLen; i++ {
		content1 := "test1_" + strconv.Itoa(i)
		content2 := "test2_" + strconv.Itoa(i)
		g.AddAttribute(i, content1)
		g.AddAttribute(i, content2)

		// success
		g.RemoveAttribute(i, 0)
		assert.Equal(t, 1, g.GetAttributesLen()[i])
		att_content := g.GetDrawData().(drawdata.Gadget).Attributes[i][0].Content
		assert.Equal(t, content2, att_content)

		// remove out of index
		assert.Error(t, g.RemoveAttribute(-1, i))
		assert.Error(t, g.RemoveAttribute(1, i))
	}

	// invalid section
	assert.Error(t, g.RemoveAttribute(-1, 0))
	assert.Error(t, g.RemoveAttribute(sectionLen, 0))
}

func TestGetDrawData(t *testing.T) {
	L := drawdata.LineWidth
	M := drawdata.Margin

	gadget, err := NewGadget(Class, utils.Point{X: 1, Y: 1})
	assert.NoError(t, err)
	assert.NotNil(t, gadget)

	data := gadget.GetDrawData()
	assert.NotNil(t, data)
	gdd := data.(drawdata.Gadget)
	assert.NotNil(t, gdd)
	assert.Equal(t, 1, gdd.X)
	assert.Equal(t, 1, gdd.Y)
	assert.Equal(t, 0, gdd.Layer)
	assert.Equal(t, drawdata.DefaultGadgetColor, gdd.Color)

	// check default att
	h := L
	maxAttWidth := 0
	for i := 0; i < len(gdd.Attributes); i++ {
		for j := 0; j < len(gdd.Attributes[i]); j++ {
			if gdd.Attributes[i][j].Width > maxAttWidth {
				maxAttWidth = gdd.Attributes[i][j].Width
			}
			h += M + gdd.Attributes[i][j].Height
		}
		h += M + L
	}
	w := maxAttWidth + M*2 + L*2
	assert.Equal(t, w, gdd.Width)
	assert.Equal(t, h, gdd.Height)
	assert.Equal(t, 3, len(gdd.Attributes))
	assert.Equal(t, "Name", gdd.Attributes[0][0].Content)
	assert.Equal(t, "Attributes", gdd.Attributes[1][0].Content)
	assert.Equal(t, "Methods", gdd.Attributes[2][0].Content)
}

func TestRegisterUpdateParentDraw(t *testing.T) {
	g := newEmptyGadget(Class, utils.Point{X: 1, Y: 1})
	mp := mockParent{}
	assert.NoError(t, g.RegisterUpdateParentDraw(mp.UpdateParentDraw))
	assert.Equal(t, 0, mp.Times)

	assert.NoError(t, g.SetPoint(utils.Point{X: 2, Y: 2}))
	assert.Equal(t, 1, mp.Times)
}
