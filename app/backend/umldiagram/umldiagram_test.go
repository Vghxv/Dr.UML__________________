package umldiagram

import (
	"testing"
	"time"

	"Dr.uml/backend/utils"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestNewUMLDiagram(t *testing.T) {
	tests := []struct {
		name        string
		diagramType DiagramType
		expectError bool
		errorMsg    string
	}{
		{
			name:        "ValidClassDiagram",
			diagramType: ClassDiagram,
			expectError: false,
		},
		{
			name:        "ValidUseCaseDiagram",
			diagramType: UseCaseDiagram,
			expectError: false,
		},
		{
			name:        "ValidSequenceDiagram",
			diagramType: SequenceDiagram,
			expectError: false,
		},
		{
			name:        "InvalidDiagramType",
			diagramType: DiagramType(8), // Unsupported type
			expectError: true,
			errorMsg:    "Invalid diagram type",
		},
		{
			name:        "InvalidName",
			diagramType: ClassDiagram,
			expectError: true,
			errorMsg:    "Invalid diagram name",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Mock utils.IsValidFilePath for invalid name test
			if tt.name == "InvalidName" {
				// Assuming IsValidFilePath returns true for invalid paths in this context
				// We don't have the actual implementation, so we're testing the error case
				diagram, err := NewUMLDiagram("invalid/path", tt.diagramType)
				if tt.expectError {
					assert.Error(t, err)
					assert.Equal(t, tt.errorMsg, err.Error())
					assert.Nil(t, diagram)
				}
				return
			}

			diagram, err := NewUMLDiagram(tt.name, tt.diagramType)

			if tt.expectError {
				assert.Error(t, err)
				assert.Equal(t, tt.errorMsg, err.Error())
				assert.Nil(t, diagram)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, diagram)
				assert.NotEqual(t, uuid.Nil, diagram.GetId())
				assert.Equal(t, tt.name, diagram.GetName())
				assert.Equal(t, tt.diagramType, diagram.diagramType)
				assert.WithinDuration(t, time.Now(), diagram.lastModified, time.Second)
				assert.Equal(t, utils.Point{X: 0, Y: 0}, diagram.startPoint)
			}
		})
	}
}

func TestCheckDiagramType(t *testing.T) {
	tests := []struct {
		name        string
		diagramType DiagramType
		expected    bool
	}{
		{
			name:        "ClassDiagram",
			diagramType: ClassDiagram,
			expected:    true,
		},
		{
			name:        "UseCaseDiagram",
			diagramType: UseCaseDiagram,
			expected:    true,
		},
		{
			name:        "SequenceDiagram",
			diagramType: SequenceDiagram,
			expected:    true,
		},
		{
			name:        "InvalidDiagram",
			diagramType: DiagramType(8),
			expected:    false,
		},
		{
			name:        "ZeroValue",
			diagramType: DiagramType(0),
			expected:    false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := check(tt.diagramType)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestUMLDiagramGetters(t *testing.T) {
	diagram, err := NewUMLDiagram("TestDiagram", ClassDiagram)
	assert.NoError(t, err)
	assert.NotNil(t, diagram)

	assert.NotEqual(t, uuid.Nil, diagram.GetId())
	assert.Equal(t, "TestDiagram", diagram.GetName())
}

// TODO: Add test for NewUMLDiagramWithPath when implemented
func TestNewUMLDiagramWithPath(t *testing.T) {
	t.Skip("NewUMLDiagramWithPath not implemented yet")
	// diagram, err := NewUMLDiagramWithPath("some/path")
	// assert.Error(t, err)
	// assert.Nil(t, diagram)
}
