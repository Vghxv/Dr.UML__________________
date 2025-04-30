package umldiagram

import (
	"testing"
	"time"

	"Dr.uml/backend/component"
	"Dr.uml/backend/utils"
	"github.com/stretchr/testify/assert"
)

func TestNewUMLDiagram(t *testing.T) {
	tests := []struct {
		name        string
		inputName   string
		diagramType DiagramType
		expectError bool
		errorMsg    string
	}{
		{
			name:        "ValidClassDiagram",
			inputName:   "test1.uml",
			diagramType: ClassDiagram,
			expectError: false,
		},
		{
			name:        "ValidUseCaseDiagram",
			inputName:   "test2.uml",
			diagramType: UseCaseDiagram,
			expectError: false,
		},
		{
			name:        "ValidSequenceDiagram",
			inputName:   "test3.uml",
			diagramType: SequenceDiagram,
			expectError: false,
		},
		{
			name:        "InvalidDiagramType",
			inputName:   "test4.uml",
			diagramType: DiagramType(8),
			expectError: true,
			errorMsg:    "Invalid diagram type",
		},
		{
			name:        "InvalidDiagramType2",
			inputName:   "test5.uml",
			diagramType: DiagramType(8787),
			expectError: true,
			errorMsg:    "Invalid diagram type",
		},
		{
			name:        "InvalidName",
			inputName:   "",
			diagramType: ClassDiagram,
			expectError: true,
			errorMsg:    "Invalid diagram name",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			diagram, err := NewUMLDiagram(tt.inputName, tt.diagramType)

			if tt.expectError {
				assert.Error(t, err)
				assert.Equal(t, tt.errorMsg, err.Error())
				assert.Nil(t, diagram)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, diagram)
				assert.Equal(t, tt.inputName, diagram.GetName())
				assert.Equal(t, tt.diagramType, diagram.diagramType)
				assert.WithinDuration(t, time.Now(), diagram.lastModified, time.Second)
				assert.Equal(t, utils.Point{X: 0, Y: 0}, diagram.startPoint)
				// New assertions
				assert.Equal(t, utils.Color{R: 255, G: 255, B: 255}, diagram.backgroundColor)
				assert.NotNil(t, diagram.components)
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
			result := isValidDiagramType(tt.diagramType)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestUMLDiagramGetters(t *testing.T) {
	diagram, err := NewUMLDiagram("TestDiagram", ClassDiagram)
	assert.NoError(t, err)
	assert.NotNil(t, diagram)

	assert.Equal(t, "TestDiagram", diagram.GetName())

	comp, err := diagram.AddGadget(component.Class, utils.Point{X: 10, Y: 20})
	assert.NoError(t, err)
	assert.NotNil(t, comp)
}

func TestNewUMLDiagramWithPath(t *testing.T) {
	tests := []struct {
		name        string
		path        string
		expectError bool
		errorMsg    string
	}{
		{
			name:        "ValidPath",
			path:        "test.uml",
			expectError: false,
		},
		{
			name:        "InvalidPath",
			path:        "",
			expectError: true,
			errorMsg:    "Invalid diagram name",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			diagram, err := NewUMLDiagramWithPath(tt.path)

			if tt.expectError {
				assert.Error(t, err)
				assert.Nil(t, diagram)
				assert.Equal(t, tt.errorMsg, err.Error())
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, diagram)
				assert.Equal(t, tt.path, diagram.GetName())
				assert.Equal(t, DiagramType(ClassDiagram), diagram.diagramType)
				assert.WithinDuration(t, time.Now(), diagram.lastModified, time.Second)
			}
		})
	}
}
