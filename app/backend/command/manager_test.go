package command

import (
	"testing"
	"time"

	"Dr.uml/backend/utils/duerror"
	"github.com/stretchr/testify/assert"
)

// MockCommand implements the Command interface for testing
type MockCommand struct {
	before       time.Time
	after        time.Time
	executeErr   duerror.DUError
	unexecuteErr duerror.DUError
}

func (m *MockCommand) Execute() duerror.DUError   { return m.executeErr }
func (m *MockCommand) Unexecute() duerror.DUError { return m.unexecuteErr }
func (m *MockCommand) GetBefore() time.Time       { return m.before }
func (m *MockCommand) GetAfter() time.Time        { return m.after }

func TestManager_Execute_NilCommand(t *testing.T) {
	m := NewManager(time.Now())
	err := m.Execute(nil)
	assert.NotNil(t, err)
	assert.Equal(t, "command is nil", err.Error())
}

func TestManager_Execute_Success(t *testing.T) {
	before := time.Now()
	after := before.Add(time.Minute)
	cmd := &MockCommand{before: before, after: after}

	m := NewManager(before)
	err := m.Execute(cmd)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(m.undoStack))
	assert.Equal(t, 0, len(m.redoStack))
	assert.Equal(t, after, m.GetLastModified())
}

func TestManager_Execute_ErrorFromCommand(t *testing.T) {
	before := time.Now()
	cmd := &MockCommand{
		before:     before,
		after:      before.Add(time.Minute),
		executeErr: duerror.NewInvalidArgumentError("exec fail"),
	}

	m := NewManager(before)
	err := m.Execute(cmd)
	assert.NotNil(t, err)
	assert.Equal(t, "exec fail", err.Error())
	assert.Equal(t, 0, len(m.undoStack))
}

func TestManager_Undo_NoCommand(t *testing.T) {
	m := NewManager(time.Now())
	err := m.Undo()
	assert.NotNil(t, err)
	assert.Equal(t, "no undo command", err.Error())
}

func TestManager_Undo_Success(t *testing.T) {
	before := time.Now()
	after := before.Add(time.Minute)
	cmd := &MockCommand{before: before, after: after}

	m := NewManager(after)
	m.Execute(cmd)
	err := m.Undo()
	assert.Nil(t, err)
	assert.Equal(t, 0, len(m.undoStack))
	assert.Equal(t, 1, len(m.redoStack))
	assert.Equal(t, before, m.GetLastModified())
}

func TestManager_Undo_ErrorFromCommand(t *testing.T) {
	before := time.Now()
	after := before.Add(time.Minute)
	cmd := &MockCommand{
		before:       before,
		after:        after,
		unexecuteErr: duerror.NewInvalidArgumentError("undo fail"),
	}

	m := NewManager(after)
	m.Execute(cmd)
	err := m.Undo()
	assert.NotNil(t, err)
	assert.Equal(t, "undo fail", err.Error())
	assert.Equal(t, 0, len(m.undoStack)) // still popped
}

func TestManager_Redo_NoCommand(t *testing.T) {
	m := NewManager(time.Now())
	err := m.Redo()
	assert.NotNil(t, err)
	assert.Equal(t, "no redo command", err.Error())
}

func TestManager_Redo_Success(t *testing.T) {
	before := time.Now()
	after := before.Add(time.Minute)
	cmd := &MockCommand{before: before, after: after}

	m := NewManager(before)
	m.Execute(cmd)
	m.Undo()
	err := m.Redo()
	assert.Nil(t, err)
	assert.Equal(t, 1, len(m.undoStack))
	assert.Equal(t, 0, len(m.redoStack))
	assert.Equal(t, after, m.GetLastModified())
}

func TestManager_Redo_ErrorFromCommand(t *testing.T) {
	before := time.Now()
	after := before.Add(time.Minute)
	cmd := &MockCommand{
		before:     before,
		after:      after,
		executeErr: duerror.NewInvalidArgumentError("redo fail"),
	}

	m := NewManager(before)
	m.Execute(cmd)
	m.Undo()
	err := m.Redo()
	assert.NotNil(t, err)
	assert.Equal(t, "no redo command", err.Error())
}

func TestManager_StackLimit(t *testing.T) {
	now := time.Now()
	m := NewManager(now)

	for i := 0; i < CMD_LIMIT+5; i++ {
		cmd := &MockCommand{
			before: now.Add(time.Duration(i) * time.Minute),
			after:  now.Add(time.Duration(i+1) * time.Minute),
		}
		m.Execute(cmd)
	}

	assert.Equal(t, CMD_LIMIT, len(m.undoStack))
	assert.Equal(t, 0, len(m.redoStack))
}
