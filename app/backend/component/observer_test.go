package component

import (
	"testing"

	"Dr.uml/backend/drawdata"
	"Dr.uml/backend/utils"
	"Dr.uml/backend/utils/duerror"
	"github.com/stretchr/testify/assert"
)

// Mock gadget for testing observer pattern
func createMockGadget() *Gadget {
	return &Gadget{
		gadgetType: Class,
		point:      utils.Point{X: 10, Y: 10},
		layer:      0,
		color:      "#FF0000",
		drawData: drawdata.Gadget{
			X:      10,
			Y:      10,
			Width:  100,
			Height: 50,
		},
		observers: make(map[interface{}]func() duerror.DUError),
	}
}

// Mock association for testing observer pattern
type mockAssociation struct {
	*Association
	updateCallCount int
}

func (m *mockAssociation) updateDrawData() duerror.DUError {
	m.updateCallCount++
	return nil
}

func (m *mockAssociation) RegisterAsObserver() duerror.DUError {
	// Register this association as an observer of its parent gadgets
	if m.parents[0] != nil {
		if err := m.parents[0].AddObserver(m.Association, m.updateDrawData); err != nil {
			return err
		}
	}
	if m.parents[1] != nil && m.parents[1] != m.parents[0] {
		if err := m.parents[1].AddObserver(m.Association, m.updateDrawData); err != nil {
			return err
		}
	}
	return nil
}

func (m *mockAssociation) UnregisterAsObserver() duerror.DUError {
	// Unregister this association as an observer of its parent gadgets
	if m.parents[0] != nil {
		if err := m.parents[0].RemoveObserver(m.Association); err != nil {
			// Don't return error if observer wasn't found, just continue
		}
	}
	if m.parents[1] != nil && m.parents[1] != m.parents[0] {
		if err := m.parents[1].RemoveObserver(m.Association); err != nil {
			// Don't return error if observer wasn't found, just continue
		}
	}
	return nil
}

func createMockAssociation(parent1, parent2 *Gadget) *mockAssociation {
	return &mockAssociation{
		Association: &Association{
			assType:         Extension,
			parents:         [2]*Gadget{parent1, parent2},
			startPointRatio: [2]float64{0.5, 0.5},
			endPointRatio:   [2]float64{0.5, 0.5},
		},
		updateCallCount: 0,
	}
}

func TestObserverBasic(t *testing.T) {
	gadget := createMockGadget()

	observerCalled := false
	testObserver := func() duerror.DUError {
		observerCalled = true
		return nil
	}

	// Test observer registration
	err := gadget.AddObserver("test", testObserver)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(gadget.observers))

	// Test observer notification
	err = gadget.notifyObservers()
	assert.NoError(t, err)
	assert.True(t, observerCalled)

	// Test observer removal
	err = gadget.RemoveObserver("test")
	assert.NoError(t, err)
	assert.Equal(t, 0, len(gadget.observers))
}

func TestObserverOnPositionChange(t *testing.T) {
	gadget := createMockGadget()

	notificationCount := 0
	testObserver := func() duerror.DUError {
		notificationCount++
		return nil
	}

	err := gadget.AddObserver("test", testObserver)
	assert.NoError(t, err)

	// Change position - should trigger observer notification
	err = gadget.SetPoint(utils.Point{X: 50, Y: 50})
	assert.NoError(t, err)

	// Should have been called once for position change
	assert.Equal(t, 1, notificationCount)
}

func TestAssociationObserverRegistration(t *testing.T) {
	gadget1 := createMockGadget()
	gadget2 := createMockGadget()
	gadget2.point = utils.Point{X: 200, Y: 200}
	gadget2.drawData.X = 200
	gadget2.drawData.Y = 200

	association := createMockAssociation(gadget1, gadget2)

	// Test association observer registration
	err := association.RegisterAsObserver()
	assert.NoError(t, err)

	// Check that both parents have the association as an observer
	_, exists1 := gadget1.observers[association.Association]
	assert.True(t, exists1)
	_, exists2 := gadget2.observers[association.Association]
	assert.True(t, exists2)

	// Test association observer unregistration
	err = association.UnregisterAsObserver()
	assert.NoError(t, err)

	// Check that both parents no longer have the association as an observer
	_, exists1 = gadget1.observers[association.Association]
	assert.False(t, exists1)
	_, exists2 = gadget2.observers[association.Association]
	assert.False(t, exists2)
}

func TestObserverPatternIntegration(t *testing.T) {
	// Create gadgets
	gadget1 := createMockGadget()
	gadget2 := createMockGadget()
	gadget2.point = utils.Point{X: 200, Y: 200}
	gadget2.drawData.X = 200
	gadget2.drawData.Y = 200

	// Create mock association that tracks update calls
	association := createMockAssociation(gadget1, gadget2)

	// Register association as observer
	err := association.RegisterAsObserver()
	assert.NoError(t, err)

	// Move gadget1 - should trigger association update
	err = gadget1.SetPoint(utils.Point{X: 100, Y: 100})
	assert.NoError(t, err)

	// Association should have been notified
	assert.Equal(t, 1, association.updateCallCount)

	// Move gadget2 - should also trigger association update
	err = gadget2.SetPoint(utils.Point{X: 300, Y: 300})
	assert.NoError(t, err)

	// Association should have been notified again
	assert.Equal(t, 2, association.updateCallCount)
}

func TestObserverErrorHandling(t *testing.T) {
	gadget := createMockGadget()

	// Test adding nil observer function
	err := gadget.AddObserver("test", nil)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "observer function is nil")

	// Test adding observer with nil key
	err = gadget.AddObserver(nil, func() duerror.DUError { return nil })
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "observer key is nil")

	// Test removing non-existent observer
	err = gadget.RemoveObserver("non-existent")
	assert.NoError(t, err) // Should not error

	// Test removing observer with nil key
	err = gadget.RemoveObserver(nil)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "observer key is nil")
}
